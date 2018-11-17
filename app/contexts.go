// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "common services": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/ZJU-DistributedAI/Common/design
// --out=$(GOPATH)/src/github.com/ZJU-DistributedAI/Common
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// AddStorageContext provides the storage add action context.
type AddStorageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *FilePayload
}

// NewAddStorageContext parses the incoming request URL and body, performs validations and creates the
// context used by the storage controller add action.
func NewAddStorageContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddStorageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddStorageContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddStorageContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CatStorageContext provides the storage cat action context.
type CatStorageContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Address string
}

// NewCatStorageContext parses the incoming request URL and body, performs validations and creates the
// context used by the storage controller cat action.
func NewCatStorageContext(ctx context.Context, r *http.Request, service *goa.Service) (*CatStorageContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CatStorageContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAddress := req.Params["address"]
	if len(paramAddress) > 0 {
		rawAddress := paramAddress[0]
		rctx.Address = rawAddress
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CatStorageContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}