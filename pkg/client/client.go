package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultBaseURL        = "https://api.openai.com/v1"
	Seconds30             = 30 * time.Second
	MaxIdleConns          = 100
	IdleConnTimeout       = 90 * time.Second
	TLSHandshakeTimeout   = 10 * time.Second
	ExpectContinueTimeout = 1 * time.Second
)

type Client struct {
	baseURL    string
	underlying *http.Client
	apiKey     string
}

func New(apiKey string, options ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		underlying: &http.Client{
			Transport: &transport{
				underlying: http.Transport{
					Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   Seconds30,
						KeepAlive: Seconds30,
					}).DialContext,
					ForceAttemptHTTP2:     true,
					MaxIdleConns:          MaxIdleConns,
					IdleConnTimeout:       IdleConnTimeout,
					TLSHandshakeTimeout:   TLSHandshakeTimeout,
					ExpectContinueTimeout: ExpectContinueTimeout,
				},
			},
		},
		apiKey: apiKey,
	}
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Client) url(path string) string {
	if strings.Contains(path, "://") {
		return path
	}

	return c.baseURL + path
}

func (c *Client) do(ctx context.Context, method, path string, target interface{}, body io.Reader) error {
	req, err := http.NewRequestWithContext(ctx, method, c.url(path), body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.underlying.Do(req)

	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if target == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) RawGet(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url(path), nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.underlying.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}

func (c *Client) Get(ctx context.Context, path string, target interface{}) error {
	return c.do(ctx, http.MethodGet, path, target, nil)
}

func (c *Client) Post(ctx context.Context, path string, target interface{}, body interface{}) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}

	return c.do(ctx, http.MethodPost, path, target, buffer)
}

func (c *Client) Delete(ctx context.Context, path string, target interface{}) error {
	return c.do(ctx, http.MethodDelete, path, target, nil)
}

func (c *Client) Patch(ctx context.Context, path string, target interface{}, body interface{}) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}

	return c.do(ctx, http.MethodPatch, path, target, buffer)
}
