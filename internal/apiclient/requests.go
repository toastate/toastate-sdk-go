package apiclient

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/toastate/toastate-sdk-go/common/models"
)

func (c *Client) request(authed bool, url, method string, body interface{}, resp interface{}) (*Error, error) {
	url = c.prepareURL(url)

	var req *http.Request
	var err error
	if body != nil {
		b := new(bytes.Buffer)
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(false)
		err = enc.Encode(body)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, b)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}

	c.setupRequest(req, authed)

	response, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	b, _ := io.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode != 200 {
		e := &Error{
			Status: response.StatusCode,
		}
		if len(b) == 0 {
			e.Code = "unhandled"
			e.Message = "The remote API did not provide any error message"
			return e, nil
		}

		err = json.Unmarshal(b, e)
		if err != nil {
			e.Code = "unhandled"
			e.Message = "The remote API provided the following invalid JSON error: " + string(b)
			return e, nil
		}

		return e, nil
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) requestStreamRawResponse(authed bool, url, method string, body interface{}) (io.ReadCloser, *Error, error) {
	url = c.prepareURL(url)

	var req *http.Request
	var err error
	if body != nil {
		b := new(bytes.Buffer)
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(false)
		err = enc.Encode(body)
		if err != nil {
			return nil, nil, err
		}

		req, err = http.NewRequest(method, url, b)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, nil, err
	}

	c.setupRequest(req, authed)

	response, err := c.http.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != 200 {
		b, _ := io.ReadAll(response.Body)
		response.Body.Close()

		e := &Error{
			Status: response.StatusCode,
		}
		if len(b) == 0 {
			e.Code = "unhandled"
			e.Message = "The remote API did not provide any error message"
			return nil, e, nil
		}

		err = json.Unmarshal(b, e)
		if err != nil {
			e.Code = "unhandled"
			e.Message = "The remote API provided the following invalid JSON error: " + string(b)
			return nil, e, nil
		}

		return nil, e, nil
	}

	return response.Body, nil, nil
}

func (c *Client) requestMultipartFolder(authed bool, folder, url, method string, body interface{}, resp interface{}) (*Error, error) {
	url = c.prepareURL(url)
	bod, err := marshalBody(body)
	if err != nil {
		return nil, err
	}

	// Create a pipe for writing from the file and reading to
	// the request concurrently.
	bodyReader, bodyWriter := io.Pipe()
	formWriter := multipart.NewWriter(bodyWriter)

	// Store the first write error in writeErr.
	var (
		writeErr error
		errOnce  sync.Once
	)
	setErr := func(err error) {
		if err != nil {
			errOnce.Do(func() { writeErr = err })
		}
	}
	go func() {
		err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			rel, err := filepath.Rel(folder, path)
			if err != nil {
				return err
			}

			partWriter, err := formWriter.CreateFormFile("file", base32.StdEncoding.EncodeToString([]byte(rel)))
			if err != nil {
				return err
			}

			f, err := os.OpenFile(path, os.O_RDONLY, 0644)
			if err != nil {
				return err
			}

			// Reduce number of syscalls when reading from disk.
			bufferedFileReader := bufio.NewReader(f)
			defer f.Close()

			for {
				n, err := io.CopyN(partWriter, bufferedFileReader, 1024*1024*5)
				if err != nil {
					if err != io.EOF {
						return err
					}
					break
				}
				if n == 0 {
					break
				}
			}

			return nil
		})

		setErr(err)

		if err == nil && len(bod) > 0 {
			err = formWriter.WriteField("request", string(bod))
			if err != nil {
				setErr(err)
			}
		}

		setErr(formWriter.Close())
		setErr(bodyWriter.Close())
	}()

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	c.setupRequest(req, authed)
	req.Header.Set("Content-Type", formWriter.FormDataContentType())

	// This operation will block until both the formWriter
	// and bodyWriter have been closed by the goroutine,
	// or in the event of a HTTP error.
	response, err := (&http.Client{
		Timeout: 3600 * time.Hour,
	}).Do(req)
	if err != nil {
		return nil, err
	}

	if writeErr != nil {
		return nil, writeErr
	}

	b, _ := io.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode != 200 {
		e := &Error{
			Status: response.StatusCode,
		}
		if len(b) == 0 {
			e.Code = "unhandled"
			e.Message = "The remote API did not provide any error message"
			return e, nil
		}

		err = json.Unmarshal(b, e)
		if err != nil {
			e.Code = "unhandled"
			e.Message = "The remote API provided the following invalid JSON error: " + string(b)
			return e, nil
		}

		return e, nil
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) requestMultipartReaders(authed bool, ch chan *models.MultipartItem, url, method string, body interface{}, resp interface{}) (*Error, error) {
	url = c.prepareURL(url)
	bod, err := marshalBody(body)
	if err != nil {
		return nil, err
	}

	// Create a pipe for writing from the file and reading to
	// the request concurrently.
	bodyReader, bodyWriter := io.Pipe()
	formWriter := multipart.NewWriter(bodyWriter)

	// Store the first write error in writeErr.
	var (
		writeErr error
		errOnce  sync.Once
	)
	setErr := func(err error) {
		if err != nil {
			errOnce.Do(func() { writeErr = err })
		}
	}
	go func() {
		var err error
		var partWriter io.Writer
		var n int64
	F1:
		for {
			item := <-ch
			if item == nil {
				break F1
			}

			partWriter, err = formWriter.CreateFormFile("file", base32.StdEncoding.EncodeToString([]byte(item.Filename)))
			if err != nil {
				item.R.Close()
				setErr(err)
				break F1
			}

			for {
				n, err = io.CopyN(partWriter, item.R, 1024*1024*5)
				if err != nil {
					if err != io.EOF {
						item.R.Close()
						setErr(err)
						break F1
					}
					break
				}
				if n == 0 {
					break
				}
			}

			item.R.Close()
		}

		if err == nil && len(bod) > 0 {
			err = formWriter.WriteField("request", string(bod))
			if err != nil {
				setErr(err)
			}
		}

		setErr(formWriter.Close())
		setErr(bodyWriter.Close())
	}()

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	c.setupRequest(req, authed)
	req.Header.Set("Content-Type", formWriter.FormDataContentType())

	// This operation will block until both the formWriter
	// and bodyWriter have been closed by the goroutine,
	// or in the event of a HTTP error.
	response, err := (&http.Client{
		Timeout: 3600 * time.Hour,
	}).Do(req)
	if err != nil {
		fmt.Println("3", err)
		return nil, err
	}

	if writeErr != nil {
		return nil, writeErr
	}

	b, _ := io.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode != 200 {
		e := &Error{
			Status: response.StatusCode,
		}
		if len(b) == 0 {
			e.Code = "unhandled"
			e.Message = "The remote API did not provide any error message"
			return e, nil
		}

		err = json.Unmarshal(b, e)
		if err != nil {
			e.Code = "unhandled"
			e.Message = "The remote API provided the following invalid JSON error: " + string(b)
			return e, nil
		}

		return e, nil
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
