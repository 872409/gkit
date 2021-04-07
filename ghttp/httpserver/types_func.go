package httpserver

import (
	"github.com/872409/gatom/gc"
)

type HTTPMethod string

const (
	HTTPPOST   HTTPMethod = "POST"
	HTTPGET    HTTPMethod = "GET"
	HTTPPUT    HTTPMethod = "PUT"
	HTTPDELETE HTTPMethod = "DELETE"
)

type EndpointHandlerFunc func(g *gc.GContext, request interface{}) (interface{}, error)
type NewRequestFunc func() interface{}

type DecodeRequestFunc func(g *gc.GContext) (interface{}, error)
type EncodeResponseFunc func(g *gc.GContext, response interface{}, err error) (interface{}, error)
type DecodeRequest interface {
}

type RenderResponseFunc func(g *gc.GContext, response interface{}, err error)
