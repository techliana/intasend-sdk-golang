package intasend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// HTTPMethod represents HTTP methods
type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

// RequestOptions contains options for making HTTP requests
type RequestOptions struct {
	Method      HTTPMethod
	Endpoint    string
	Body        interface{}
	QueryParams map[string]string
	Headers     map[string]string
	UseToken    bool // Whether to use Bearer token authentication
	UseAPIKey   bool // Whether to use API key authentication
}

// HTTPResponse represents the response from an HTTP request
type HTTPResponse struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

// DoRequest performs an HTTP request with the given options
func (c *Client) DoRequest(opts *RequestOptions) (*HTTPResponse, error) {
	// Build the full URL
	fullURL, err := c.buildURL(opts.Endpoint, opts.QueryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	// Prepare request body
	var bodyReader io.Reader
	if opts.Body != nil {
		bodyBytes, err := c.prepareBody(opts.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to prepare request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)

		// Log request body if debugging is enabled
		if c.shouldLog() {
			fmt.Printf("Request Body: %s\n", string(bodyBytes))
		}
	}

	// Create HTTP request
	req, err := http.NewRequest(string(opts.Method), fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	c.setHeaders(req, opts)

	// Log request details if debugging is enabled
	if c.shouldLog() {
		c.logRequest(req)
	}

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	httpResp := &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}

	// Log response if debugging is enabled
	if c.shouldLog() {
		c.logResponse(httpResp)
	}

	return httpResp, nil
}

// DoRequestWithJSON performs an HTTP request and unmarshals the response into the provided interface
func (c *Client) DoRequestWithJSON(opts *RequestOptions, result interface{}) error {
	resp, err := c.DoRequest(opts)
	if err != nil {
		return err
	}

	// Handle error responses
	if resp.StatusCode >= 400 {
		return c.handleErrorResponse(resp)
	}

	// Unmarshal successful response
	if result != nil {
		if err := json.Unmarshal(resp.Body, result); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
	}

	return nil
}

// buildURL constructs the full URL with query parameters
func (c *Client) buildURL(endpoint string, queryParams map[string]string) (string, error) {
	// If endpoint is already a full URL, use it as is
	var fullURL string
	if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		fullURL = endpoint
	} else {
		// Remove leading slash if present to avoid double slashes
		endpoint = strings.TrimPrefix(endpoint, "/")
		fullURL = fmt.Sprintf("%s/%s", strings.TrimSuffix(c.BaseURL, "/"), endpoint)
	}

	if len(queryParams) == 0 {
		return fullURL, nil
	}

	u, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// prepareBody converts the body interface to bytes
func (c *Client) prepareBody(body interface{}) ([]byte, error) {
	switch v := body.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	case io.Reader:
		return io.ReadAll(v)
	default:
		// Assume it's a struct that should be JSON marshaled
		return json.Marshal(v)
	}
}

// setHeaders sets the appropriate headers for the request
func (c *Client) setHeaders(req *http.Request, opts *RequestOptions) {
	// Set default headers
	req.Header.Set("Accept", "application/json")

	// Set Content-Type for requests with body
	if opts.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Set authentication headers
	if opts.UseToken && c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	if opts.UseAPIKey && c.PublishableKey != "" {
		// Different endpoints use different header names for API key
		req.Header.Set("X-IntaSend-Public-API-Key", c.PublishableKey)
		req.Header.Set("X-IntaSend-Public-Key-Id", c.PublishableKey)
	}

	// Set custom headers
	for key, value := range opts.Headers {
		req.Header.Set(key, value)
	}
}

// handleErrorResponse processes error responses and returns appropriate errors
func (c *Client) handleErrorResponse(resp *HTTPResponse) error {
	var errorResp ErrorResponse
	if err := json.Unmarshal(resp.Body, &errorResp); err == nil && errorResp.Message != "" {
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, errorResp.Message)
	}
	return fmt.Errorf("API error (%d): %s", resp.StatusCode, string(resp.Body))
}

// shouldLog determines if logging should be enabled
// You can extend this to check environment variables or client settings
func (c *Client) shouldLog() bool {
	// For now, always log. You can make this configurable later
	return c.ShowLogs
}

// logRequest logs the HTTP request details
func (c *Client) logRequest(req *http.Request) {
	fmt.Printf("=== HTTP REQUEST ===\n")
	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("URL: %s\n", req.URL.String())
	fmt.Printf("Headers:\n")
	for key, values := range req.Header {
		for _, value := range values {
			// Mask sensitive headers
			if strings.ToLower(key) == "authorization" {
				fmt.Printf("  %s: %s\n", key, maskToken(value))
			} else {
				fmt.Printf("  %s: %s\n", key, value)
			}
		}
	}
	fmt.Printf("==================\n\n")
}

// logResponse logs the HTTP response details
func (c *Client) logResponse(resp *HTTPResponse) {
	fmt.Printf("=== HTTP RESPONSE ===\n")
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(resp.Body))
	fmt.Printf("===================\n\n")
}

// maskToken masks sensitive token information for logging
func maskToken(token string) string {
	if len(token) <= 10 {
		return "***"
	}
	return token[:4] + "***" + token[len(token)-4:]
}

// Convenience methods for common HTTP operations

// Get performs a GET request
func (c *Client) Get(endpoint string, queryParams map[string]string, useToken bool) (*HTTPResponse, error) {
	return c.DoRequest(&RequestOptions{
		Method:      GET,
		Endpoint:    endpoint,
		QueryParams: queryParams,
		UseToken:    useToken,
	})
}

// Post performs a POST request
func (c *Client) Post(endpoint string, body interface{}, useToken, useAPIKey bool) (*HTTPResponse, error) {
	return c.DoRequest(&RequestOptions{
		Method:    POST,
		Endpoint:  endpoint,
		Body:      body,
		UseToken:  useToken,
		UseAPIKey: useAPIKey,
	})
}

// Put performs a PUT request
func (c *Client) Put(endpoint string, body interface{}, useToken bool) (*HTTPResponse, error) {
	return c.DoRequest(&RequestOptions{
		Method:   PUT,
		Endpoint: endpoint,
		Body:     body,
		UseToken: useToken,
	})
}

// Delete performs a DELETE request
func (c *Client) Delete(endpoint string, useToken bool) (*HTTPResponse, error) {
	return c.DoRequest(&RequestOptions{
		Method:   DELETE,
		Endpoint: endpoint,
		UseToken: useToken,
	})
}

// GetJSON performs a GET request and unmarshals the response
func (c *Client) GetJSON(endpoint string, queryParams map[string]string, useToken bool, result interface{}) error {
	return c.DoRequestWithJSON(&RequestOptions{
		Method:      GET,
		Endpoint:    endpoint,
		QueryParams: queryParams,
		UseToken:    useToken,
	}, result)
}

// PostJSON performs a POST request and unmarshals the response
func (c *Client) PostJSON(endpoint string, body interface{}, useToken, useAPIKey bool, result interface{}) error {
	return c.DoRequestWithJSON(&RequestOptions{
		Method:    POST,
		Endpoint:  endpoint,
		Body:      body,
		UseToken:  useToken,
		UseAPIKey: useAPIKey,
	}, result)
}
