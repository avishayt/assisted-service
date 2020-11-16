// Code generated by go-swagger; DO NOT EDIT.

package bootfiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DownloadBootFilesHandlerFunc turns a function with the right signature into a download boot files handler
type DownloadBootFilesHandlerFunc func(DownloadBootFilesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadBootFilesHandlerFunc) Handle(params DownloadBootFilesParams) middleware.Responder {
	return fn(params)
}

// DownloadBootFilesHandler interface for that can handle valid download boot files params
type DownloadBootFilesHandler interface {
	Handle(DownloadBootFilesParams) middleware.Responder
}

// NewDownloadBootFiles creates a new http.Handler for the download boot files operation
func NewDownloadBootFiles(ctx *middleware.Context, handler DownloadBootFilesHandler) *DownloadBootFiles {
	return &DownloadBootFiles{Context: ctx, Handler: handler}
}

/*DownloadBootFiles swagger:route GET /boot-files bootfiles downloadBootFiles

Downloads files used for booting servers.

*/
type DownloadBootFiles struct {
	Context *middleware.Context
	Handler DownloadBootFilesHandler
}

func (o *DownloadBootFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDownloadBootFilesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
