package http

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		contentType := strings.ToLower(c.GetHeader("Content-Type"))
		isMultipart := strings.Contains(contentType, "multipart/form-data")
		var bodyBytes []byte
		if !isMultipart && c.Request != nil && c.Request.Body != nil {
			b, _ := io.ReadAll(c.Request.Body)
			bodyBytes = b
			c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
		}

		c.Next()

		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		statusMessage := http.StatusText(statusCode)
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		if c.Request.URL.Path == "/api/v1/health-check" {
			return
		}

		headers := map[string][]string{}
		for k, v := range c.Request.Header {
			if strings.ToLower(k) == "authorization" {
				continue
			}
			headers[k] = v
		}

		entry := log.WithFields(log.Fields{
			"status":     statusCode,
			"method":     c.Request.Method,
			"path":       path,
			"raw_path":   c.Request.URL.Path,
			"query":      c.Request.URL.RawQuery,
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
			"latency_ms": float64(latency.Microseconds()) / 1000.0,
			"size_bytes": c.Writer.Size(),
			"referer":    c.Request.Referer(),
			"request_id": c.GetHeader("X-Request-ID"),
			"message":    statusMessage,
		})

		if !isMultipart {
			if len(bodyBytes) > 0 {
				entry = entry.WithField("body", string(bodyBytes))
			}
			if len(headers) > 0 {
				entry = entry.WithField("headers", headers)
			}
		}

		switch {
		case statusCode >= 500:
			entry.Error("HTTP request completed")
		case statusCode >= 400:
			entry.Warn("HTTP request completed")
		default:
			entry.Info("HTTP request completed")
		}
	}
}
