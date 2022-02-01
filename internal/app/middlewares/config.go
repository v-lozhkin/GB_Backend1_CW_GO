package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/v-lozhkin/GB_Backend1_CW_GO/internal/config"
	contextUtils "github.com/v-lozhkin/GB_Backend1_CW_GO/internal/pkg/context"
)

func ConfigMiddleware(cfg config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			newCtx := contextUtils.SetConfig(ectx.Request().Context(), cfg)
			ectx.SetRequest(ectx.Request().WithContext(newCtx))

			return next(ectx)
		}
	}
}
