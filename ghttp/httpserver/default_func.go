package httpserver

import (
	"github.com/872409/gatom/gc"
	"github.com/872409/gatom/log"
	"github.com/872409/gatom/util"
)

func handleBindJSON(g *gc.GContext, newRequest NewRequestFunc) (interface{}, error) {
	request := newRequest()
	err := g.BindJSONWithError(request)
	log.Errorln("handleBindJSON:", err)
	return request, err
}

func decodeRequestBindJSON(g *gc.GContext, newRequest interface{}) error {
	return g.BindJSONWithError(newRequest)
}

func renderJSON(g *gc.GContext, response interface{}, err error) {
	if err != nil {
		coderError, ok := err.(util.CodeError)
		if ok {
			g.JSONCodeError(coderError)
		} else {
			g.JSONError(err.Error())
		}
		return
	}
	g.JSONSuccess(response)
}
