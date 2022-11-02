package plexnet

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

// GinZeroLogger Wraps the zerologger into a Gin Middleware
//so that the standard gin logs also appear in our desired format
func GinZeroLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		done := time.Since(start)

		errorMsg := c.Errors.ByType(gin.ErrorTypePrivate).String()

		logEvent := log.Info()
		if errorMsg != "" {
			logEvent = log.Error()
		}

		logEvent.Int("statusCode", c.Writer.Status()).
			Dur("duration", done).
			Str("clientIP", c.ClientIP()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Msg(errorMsg)
	}
}

func RequestIDLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			TODO
			What needs to be done here:
			Step 1
			- Read Header request ID
			- Create new if not existent
			- Create logger instance with contextual fields from headers
			- Save logger in the gin context with c.Set(key, value)

			Step 2 (In Handler Function)
			- get Logger from gin context
			- create a GO Value Context and store the logger
			- pass context to all functions
			- when logging, always log from context instead of global
			- Can be made way easier with some helper functions /packages
		*/
	}
}
