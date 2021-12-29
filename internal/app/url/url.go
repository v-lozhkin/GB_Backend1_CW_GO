package url

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func Create(ectx echo.Context) error {
	ectx.Logger().Info("Create")
	link := &models.Link{}
	if err := ectx.Bind(link); err != nil {
		return err
	}

	return ectx.JSON(http.StatusOK, link)
}
