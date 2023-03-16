package mlog

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Middleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return logMiddleware(next, logger)
	}
}

func logMiddleware(next echo.HandlerFunc, logger *zap.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		logstash(c, logger, next) // POC
		l := logParentID(c, logger)
		c.Set(key, l)
		return next(c)
	}
}

func logParentID(c echo.Context, logger *zap.Logger) *zap.Logger {
	xParent := c.Request().Header.Get("X-Parent-ID")
	if xParent == "" {
		xParent = uuid.NewString()
	}
	xSpan := uuid.NewString()
	return logger.With(zap.String("parent-id", xParent),
		zap.String("span-id", xSpan))
}

// log the request print into terminal
func logstash(c echo.Context, logger *zap.Logger, next echo.HandlerFunc) {
	start := time.Now()

	err := next(c)
	if err != nil {
		c.Error(err)
	}

	req := c.Request()
	res := c.Response()

	fields := []zapcore.Field{
		zap.String("remote_ip", c.RealIP()),
		zap.String("latency", time.Since(start).String()),
		zap.String("host", req.Host),
		zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
		zap.Int("status", res.Status),
		zap.Int64("size", res.Size),
		zap.String("user_agent", req.UserAgent()),
	}

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	fields = append(fields, zap.String("request_id", id))

	n := res.Status
	switch {
	case n >= 500:
		logger.With(zap.Error(err)).Error("Server error", fields...)
	case n >= 400:
		logger.With(zap.Error(err)).Warn("Client error", fields...)
	case n >= 300:
		logger.Info("Redirection", fields...)
	default:
		logger.Info("Success", fields...)
	}
}