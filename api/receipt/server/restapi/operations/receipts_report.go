// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReceiptsReportHandlerFunc turns a function with the right signature into a receipts report handler
type ReceiptsReportHandlerFunc func(ReceiptsReportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReceiptsReportHandlerFunc) Handle(params ReceiptsReportParams) middleware.Responder {
	return fn(params)
}

// ReceiptsReportHandler interface for that can handle valid receipts report params
type ReceiptsReportHandler interface {
	Handle(ReceiptsReportParams) middleware.Responder
}

// NewReceiptsReport creates a new http.Handler for the receipts report operation
func NewReceiptsReport(ctx *middleware.Context, handler ReceiptsReportHandler) *ReceiptsReport {
	return &ReceiptsReport{Context: ctx, Handler: handler}
}

/*ReceiptsReport swagger:route POST /receipts_report receiptsReport

Receipt report

*/
type ReceiptsReport struct {
	Context *middleware.Context
	Handler ReceiptsReportHandler
}

func (o *ReceiptsReport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReceiptsReportParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
