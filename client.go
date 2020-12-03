package goGrowiAPI

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Config struct {
	URL   string
	Token string
}

type Client struct {
	httpClient *http.Client
	config     *Config

	Pages *PagesService
}

type service struct {
	client *Client
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.URL == "" {
		return nil, fmt.Errorf("missing growi url")
	}

	if cfg.Token == "" {
		return nil, fmt.Errorf("missing api token")
	}

	httpClient := http.DefaultClient

	client := Client{
		httpClient: httpClient,
		config:     &cfg,
	}
	client.Pages = &PagesService{client: &client}
	return &client, nil
}

func isMethod(method string) bool {
	for _, httpMethod := range []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut} {
		if method == httpMethod {
			return true
		}
	}
	return false
}

func (c *Client) newRequest(ctx context.Context, method string, uri string, params *url.Values) ([]byte, error) {
	if !isMethod(method) {
		return nil, fmt.Errorf("failed not http metod %s", method)
	}

	u, err := url.Parse(c.config.URL)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if method == http.MethodGet {
		u.RawQuery = params.Encode()
	} else {
		body = strings.NewReader(params.Encode())
	}

	u.Path = path.Join(u.Path, uri)

	req, err := http.NewRequest(method, u.String(), body)
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
