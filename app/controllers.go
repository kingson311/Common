// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "common services": Application Controllers
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

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// SchemaController is the controller interface for the Schema actions.
type SchemaController interface {
	goa.Muxer
	goa.FileServer
}

// MountSchemaController "mounts" a Schema resource controller on the given service.
func MountSchemaController(service *goa.Service, ctrl SchemaController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/schema/*filepath", "public/schema/")
	service.Mux.Handle("GET", "/schema/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Schema", "files", "public/schema/", "route", "GET /schema/*filepath")

	h = ctrl.FileHandler("/schema/", "public/schema/index.html")
	service.Mux.Handle("GET", "/schema/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Schema", "files", "public/schema/index.html", "route", "GET /schema/")
}

// StorageController is the controller interface for the Storage actions.
type StorageController interface {
	goa.Muxer
	Add(*AddStorageContext) error
	Cat(*CatStorageContext) error
}

// MountStorageController "mounts" a Storage resource controller on the given service.
func MountStorageController(service *goa.Service, ctrl StorageController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddStorageContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*FilePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("POST", "/storage", ctrl.MuxHandler("add", h, unmarshalAddStoragePayload))
	service.LogInfo("mount", "ctrl", "Storage", "action", "Add", "route", "POST /storage")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCatStorageContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Cat(rctx)
	}
	service.Mux.Handle("GET", "/storage/:address", ctrl.MuxHandler("cat", h, nil))
	service.LogInfo("mount", "ctrl", "Storage", "action", "Cat", "route", "GET /storage/:address")
}

// unmarshalAddStoragePayload unmarshals the request body into the context request data Payload field.
func unmarshalAddStoragePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	var err error
	var payload filePayload
	_, rawFile, err2 := req.FormFile("file")
	if err2 == nil {
		payload.File = rawFile
	} else {
		err = goa.MergeErrors(err, goa.InvalidParamTypeError("file", "file", "file"))
	}
	if err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger/*filepath", "public/swagger/")
	service.Mux.Handle("GET", "/swagger/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/", "route", "GET /swagger/*filepath")

	h = ctrl.FileHandler("/swagger/", "public/swagger/index.html")
	service.Mux.Handle("GET", "/swagger/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/index.html", "route", "GET /swagger/")
}
