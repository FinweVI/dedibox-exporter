package online

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// API Base URL
const API = "https://api.online.net/api/v1"

// Option is a functional option for configuring the API client
type Option func(*Client) error

// BaseURL allows overriding of API client baseURL for testing
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// AuthToken allows setting the API client AuthToken for querying the API
func AuthToken(token string) Option {
	return func(c *Client) error {
		c.authToken = token
		return nil
	}
}

// parseOptions parses the supplied options functions and returns a configured
// *Client instance
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// Client holds information necessary to make a request to your API
type Client struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

// NewClient creates a new API client
func NewClient(opts ...Option) (*Client, error) {
	client := &Client{
		baseURL: API,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	if err := client.parseOptions(opts...); err != nil {
		return nil, err
	}

	if client.authToken == "" {
		return nil, fmt.Errorf("AuthToken needs to be set")
	}

	return client, nil
}

func (c Client) query(urlPart string) ([]byte, error) {
	var output []byte

	u, err := url.Parse(fmt.Sprintf("%s/%s", c.baseURL, urlPart))
	if err != nil {
		return output, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return output, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return output, fmt.Errorf("invalid HTTP status code returned: %s", strconv.Itoa(resp.StatusCode))
	}

	output, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}

	return output, nil
}
