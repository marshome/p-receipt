// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ReceiptsExtract receipts extract
*/
func (a *Client) ReceiptsExtract(params *ReceiptsExtractParams) (*ReceiptsExtractOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReceiptsExtractParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "receipts_extract",
		Method:             "POST",
		PathPattern:        "/receipts_extract",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReceiptsExtractReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReceiptsExtractOK), nil

}

/*
ReceiptsReport receipts report
*/
func (a *Client) ReceiptsReport(params *ReceiptsReportParams) (*ReceiptsReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReceiptsReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "receipts_report",
		Method:             "POST",
		PathPattern:        "/receipts_report",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReceiptsReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReceiptsReportOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
