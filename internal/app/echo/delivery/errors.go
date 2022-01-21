package delivery

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Status  int         `json:"status"`
}

// Error makes it compatible with `error` interface.
func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v, status=%v", he.Code, he.Message, he.Status)
}
func HTTPErrorHandler(err error, ectx echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		j, ok := err.(*jwt.ValidationError)
		if ok {
			he = &echo.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: j.Error(),
			}
		} else {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}
	e := &HTTPError{
		Message: he.Message,
		Code:    0,
		Status:  he.Code,
	}
	err = ectx.JSON(he.Code, e)
	if err != nil {
		ectx.Echo().Logger.Error(err)
	}
}
