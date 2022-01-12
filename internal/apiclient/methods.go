package apiclient

import (
	"io"

	"github.com/toastate/toastate-sdk-go/common/models"
)

func (c *Client) AuthedGet(url string, resp interface{}) (*Error, error) {
	return c.request(true, url, "GET", nil, resp)
}

func (c *Client) AuthedStreamedGet(url string) (io.ReadCloser, *Error, error) {
	return c.requestStreamRawResponse(true, url, "GET", nil)
}

func (c *Client) Get(url string, resp interface{}) (*Error, error) {
	return c.request(false, url, "GET", nil, resp)
}

func (c *Client) AuthedPost(url string, body interface{}, resp interface{}) (*Error, error) {
	return c.request(true, url, "POST", body, resp)
}

func (c *Client) Post(url string, body interface{}, resp interface{}) (*Error, error) {
	return c.request(false, url, "POST", body, resp)
}

func (c *Client) AuthedDelete(url string, body interface{}, resp interface{}) (*Error, error) {
	return c.request(true, url, "DELETE", body, resp)
}

func (c *Client) Delete(url string, body interface{}, resp interface{}) (*Error, error) {
	return c.request(false, url, "DELETE", body, resp)
}

func (c *Client) AuthedPut(url string, body interface{}, resp interface{}) (*Error, error) {
	return c.request(true, url, "PUT", body, resp)
}

func (c *Client) Put(url string, body interface{}, resp interface{}) (*Error, error) {
	return c.request(false, url, "PUT", body, resp)
}

func (c *Client) AuthedMultipartFolderPost(folder string, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartFolder(true, folder, url, "POST", body, resp)
}

func (c *Client) MultipartFolderPost(folder string, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartFolder(false, folder, url, "POST", body, resp)
}

func (c *Client) AuthedMultipartReadersPost(ch chan *models.MultipartItem, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartReaders(true, ch, url, "POST", body, resp)
}

func (c *Client) MultipartReadersPost(ch chan *models.MultipartItem, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartReaders(false, ch, url, "POST", body, resp)
}

func (c *Client) AuthedMultipartFolderPut(folder string, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartFolder(true, folder, url, "PUT", body, resp)
}

func (c *Client) MultipartFolderPut(folder string, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartFolder(false, folder, url, "PUT", body, resp)
}

func (c *Client) AuthedMultipartReadersPut(ch chan *models.MultipartItem, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartReaders(true, ch, url, "PUT", body, resp)
}

func (c *Client) MultipartReadersPut(ch chan *models.MultipartItem, url string, body interface{}, resp interface{}) (*Error, error) {
	return c.requestMultipartReaders(false, ch, url, "PUT", body, resp)
}
