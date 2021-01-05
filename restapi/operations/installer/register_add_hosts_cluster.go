// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterAddHostsClusterHandlerFunc turns a function with the right signature into a register add hosts cluster handler
type RegisterAddHostsClusterHandlerFunc func(RegisterAddHostsClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterAddHostsClusterHandlerFunc) Handle(params RegisterAddHostsClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RegisterAddHostsClusterHandler interface for that can handle valid register add hosts cluster params
type RegisterAddHostsClusterHandler interface {
	Handle(RegisterAddHostsClusterParams, interface{}) middleware.Responder
}

// NewRegisterAddHostsCluster creates a new http.Handler for the register add hosts cluster operation
func NewRegisterAddHostsCluster(ctx *middleware.Context, handler RegisterAddHostsClusterHandler) *RegisterAddHostsCluster {
	return &RegisterAddHostsCluster{Context: ctx, Handler: handler}
}

/*RegisterAddHostsCluster swagger:route POST /add_hosts_clusters installer registerAddHostsCluster

Creates a new OpenShift cluster definition for adding nodes to and existing OCP cluster.

*/
type RegisterAddHostsCluster struct {
	Context *middleware.Context
	Handler RegisterAddHostsClusterHandler
}

func (o *RegisterAddHostsCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterAddHostsClusterParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
