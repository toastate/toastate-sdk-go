package apiclient

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	domainName string
	apiVersion string

	authToken string

	http *http.Client
}

func NewClient(domainName, apiVersion string) *Client {
	if apiVersion == "" {
		panic(fmt.Errorf("api version cannot be empty"))
	}
	if domainName == "" {
		panic(fmt.Errorf("api domain name cannot be empty"))
	}

	return &Client{
		domainName: domainName,
		apiVersion: apiVersion,
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) setupRequest(req *http.Request, authed bool) {
	req.Header.Add("X-TOASTATE-APIVERSION", c.apiVersion)
	req.Header.Add("Content-Type", "application/json")

	if authed {
		req.Header.Add("X-TOASTATE-AUTH", c.authToken)
	}
}

func (c *Client) SetAuthToken(token string) *Client {
	c.authToken = token
	return c
}

func (c *Client) prepareURL(url string) string {
	if url == "" {
		panic(fmt.Errorf("empty url"))
	}
	if url[0] != '/' {
		url = "/" + url
	}
	url = "https://" + c.domainName + url
	return url
}
