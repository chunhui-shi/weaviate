//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateActionsReferencesCreateHandlerFunc turns a function with the right signature into a weaviate actions references create handler
type WeaviateActionsReferencesCreateHandlerFunc func(WeaviateActionsReferencesCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsReferencesCreateHandlerFunc) Handle(params WeaviateActionsReferencesCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsReferencesCreateHandler interface for that can handle valid weaviate actions references create params
type WeaviateActionsReferencesCreateHandler interface {
	Handle(WeaviateActionsReferencesCreateParams, *models.Principal) middleware.Responder
}

// NewWeaviateActionsReferencesCreate creates a new http.Handler for the weaviate actions references create operation
func NewWeaviateActionsReferencesCreate(ctx *middleware.Context, handler WeaviateActionsReferencesCreateHandler) *WeaviateActionsReferencesCreate {
	return &WeaviateActionsReferencesCreate{Context: ctx, Handler: handler}
}

/*WeaviateActionsReferencesCreate swagger:route POST /actions/{id}/references/{propertyName} actions weaviateActionsReferencesCreate

Add a single reference to a class-property when cardinality is set to 'hasMany'.

Add a single reference to a class-property when cardinality is set to 'hasMany'.

*/
type WeaviateActionsReferencesCreate struct {
	Context *middleware.Context
	Handler WeaviateActionsReferencesCreateHandler
}

func (o *WeaviateActionsReferencesCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsReferencesCreateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
