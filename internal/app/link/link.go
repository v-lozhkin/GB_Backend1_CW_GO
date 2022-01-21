package link

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

type Delivery interface {
	Create(ectx echo.Context) error
}

type Usecase interface {
	Create(ctx context.Context, link *models.Link) error
}

type Repository interface {
	Create(ctx context.Context, link *models.Link) error
}
