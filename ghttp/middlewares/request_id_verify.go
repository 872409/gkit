package middlewares

import (
	"github.com/872409/gatom/gc"
	"github.com/gin-gonic/gin"
)

const (
	GRequestIdKey = "g-request-id"
)
const (
	DuplicateKeyExpireSec = 60 * 24 * 30
)

type SetNotExists func(key string, value interface{}, expiredSec int) bool

func HandleDuplicateRequestWithHeaderKey(requestIdKey string, duplicateKeyCheck SetNotExists, emptyKeyErr error, duplicateKeyErr error, expiredSec int) gin.HandlerFunc {
	return HandleDuplicateRequest(func(g *gc.GContext) string {
		return g.GetGRequestId(requestIdKey)
	}, duplicateKeyCheck, emptyKeyErr, duplicateKeyErr, expiredSec)
}

type GetKey func(g *gc.GContext) string

func HandleDuplicateRequest(getKey GetKey, duplicateKeyCheck SetNotExists, emptyKeyErr error, duplicateKeyErr error, expiredSec int) gin.HandlerFunc {

	return func(ginContent *gin.Context) {
		g := gc.New(ginContent)
		uniqueKey := getKey(g)

		if len(uniqueKey) == 0 {
			g.JSONCodeError(emptyKeyErr)
			return
		}

		if duplicateKeyCheck(uniqueKey, 1, expiredSec) {
			g.JSONCodeError(duplicateKeyErr)
			return
		}
		
		ginContent.Next()
	}
}
