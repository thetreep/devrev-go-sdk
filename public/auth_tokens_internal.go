// Package public provides primitives to interact with the openapi HTTP API.
//
// This file contains manual extensions for internal API endpoints.
// It will NOT be overwritten by oapi-codegen regeneration.
package public

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Request builders -----------------------------------------------------------

// NewAuthTokensCreateInternalRequest calls the generic AuthTokensCreateInternal builder with application/json body.
// This uses the /internal/auth-tokens.create endpoint which provides better debug error messages.
func NewAuthTokensCreateInternalRequest(server string, body AuthTokensCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAuthTokensCreateInternalRequestWithBody(server, "application/json", bodyReader)
}

// NewAuthTokensCreateInternalRequestWithBody generates requests for AuthTokensCreateInternal with any type of body.
// This uses the /internal/auth-tokens.create endpoint which provides better debug error messages.
func NewAuthTokensCreateInternalRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/internal/auth-tokens.create")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// Client methods -------------------------------------------------------------

// AuthTokensCreateInternalWithBody is like AuthTokensCreateWithBody but uses the internal endpoint
// which provides better debug error messages.
func (c *Client) AuthTokensCreateInternalWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthTokensCreateInternalRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// AuthTokensCreateInternal is like AuthTokensCreate but uses the internal endpoint
// which provides better debug error messages.
func (c *Client) AuthTokensCreateInternal(ctx context.Context, body AuthTokensCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthTokensCreateInternalRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// ClientWithResponses methods ------------------------------------------------

// AuthTokensCreateInternalWithBodyWithResponse is like AuthTokensCreateWithBodyWithResponse but uses the internal endpoint
// which provides better debug error messages.
func (c *ClientWithResponses) AuthTokensCreateInternalWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthTokensCreateResponse, error) {
	client, ok := c.ClientInterface.(*Client)
	if !ok {
		return nil, fmt.Errorf("ClientWithResponses must wrap a *Client to use AuthTokensCreateInternalWithBodyWithResponse")
	}
	rsp, err := client.AuthTokensCreateInternalWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthTokensCreateResponse(rsp)
}

// AuthTokensCreateInternalWithResponse is like AuthTokensCreateWithResponse but uses the internal endpoint
// which provides better debug error messages.
func (c *ClientWithResponses) AuthTokensCreateInternalWithResponse(ctx context.Context, body AuthTokensCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*AuthTokensCreateResponse, error) {
	client, ok := c.ClientInterface.(*Client)
	if !ok {
		return nil, fmt.Errorf("ClientWithResponses must wrap a *Client to use AuthTokensCreateInternalWithResponse")
	}
	rsp, err := client.AuthTokensCreateInternal(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthTokensCreateResponse(rsp)
}
