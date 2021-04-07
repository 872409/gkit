package httpserver

import (
	"github.com/872409/gatom/gc"
	"github.com/gin-gonic/gin"
)

func NewGETEndpoint(route *gin.RouterGroup, relativePath string) *APIEndpoint {
	return NewEndpoint(route, relativePath, HTTPGET)
}

func NewPUTEndpoint(route *gin.RouterGroup, relativePath string) *APIEndpoint {
	return NewEndpoint(route, relativePath, HTTPPUT)
}

func NewEndpoint(route *gin.RouterGroup, relativePath string, httpMethods ...HTTPMethod) *APIEndpoint {

	method := HTTPPOST
	if len(httpMethods) > 0 {
		method = httpMethods[0]
	}

	endpoint := &APIEndpoint{
		RouteGroup:     route,
		httpMethod:     method,
		RelativePath:   relativePath,
		renderResponse: renderJSON,
	}

	return endpoint
}

type APIEndpoint struct {
	RouteGroup     *gin.RouterGroup
	httpMethod     HTTPMethod
	RelativePath   string
	handler        EndpointHandlerFunc
	renderResponse RenderResponseFunc
	decodeRequest  DecodeRequestFunc
	encodeResponse EncodeResponseFunc
}

func (receiver *APIEndpoint) NewRequest(newRequest NewRequestFunc) *APIEndpoint {
	receiver.decodeRequest = func(g *gc.GContext) (interface{}, error) {
		return handleBindJSON(g, newRequest)
	}
	return receiver
}

func (receiver *APIEndpoint) DecodeRequest(decodeRequest DecodeRequestFunc) *APIEndpoint {
	receiver.decodeRequest = decodeRequest
	return receiver
}

func (receiver *APIEndpoint) HttpMethod(httpMethod HTTPMethod) *APIEndpoint {
	receiver.httpMethod = httpMethod
	return receiver
}

func (receiver *APIEndpoint) EncodeResponse(encodeResponse EncodeResponseFunc) *APIEndpoint {
	receiver.encodeResponse = encodeResponse
	return receiver
}

func (receiver *APIEndpoint) ResponseRender(renderResponse RenderResponseFunc) *APIEndpoint {
	receiver.renderResponse = renderResponse
	return receiver
}

func (receiver *APIEndpoint) Handle(handler EndpointHandlerFunc) {
	receiver.handler = handler
	work(receiver)
}

func work(endpoint *APIEndpoint) {
	endpoint.RouteGroup.Handle(string(endpoint.httpMethod), endpoint.RelativePath, func(context *gin.Context) {
		g := gc.New(context)

		var request interface{}
		if endpoint.decodeRequest != nil {
			decodeRequest, err := endpoint.decodeRequest(g)
			if err != nil {
				endpoint.renderResponse(g, nil, err)
				return
			}
			request = decodeRequest
		}

		response, err := endpoint.handler(g, request)
		if endpoint.encodeResponse != nil {
			response, err = endpoint.encodeResponse(g, response, err)
		}

		endpoint.renderResponse(g, response, err)
	})
}
