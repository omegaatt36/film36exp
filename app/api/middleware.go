package api

import (
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/omegaatt36/film36exp/logging"
	"go.uber.org/zap"
)

func injectLogging(skipPaths []string) gin.HandlerFunc {
	mSkipPaths := make(map[string]struct{})
	for _, path := range skipPaths {
		mSkipPaths[path] = struct{}{}
	}

	return func(c *gin.Context) {
		ctxWithLogger := logging.NewContext(c.Request.Context())
		c.Request = c.Request.WithContext(ctxWithLogger)

		path := c.Request.URL.Path
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		if _, ok := mSkipPaths[path]; ok {
			return
		}

		fnLog := logging.InfoWithFieldCtx

		status := c.Writer.Status()
		switch {
		case status >= 500:
			fnLog = logging.ErrorWithFieldCtx
		case status >= 400:
			fnLog = logging.WarnWithFieldCtx
		}

		fnLog(
			c.Request.Context(),
			path,
			zap.Any("data", map[string]any{
				"status":   status,
				"method":   c.Request.Method,
				"fullPath": c.FullPath(),
				"latency":  latency.Milliseconds(),
			}),
		)
	}
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logging.ErrorWithFieldCtx(
					c.Request.Context(),
					"[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				c.JSON(http.StatusInternalServerError, "recover from panic")

				// Discontinue the request handler chain processing.
				c.Abort()
			}
		}()

		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	config.AllowAllOrigins = true

	return cors.New(config)
}
