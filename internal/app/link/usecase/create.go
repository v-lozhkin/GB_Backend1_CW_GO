package usecase

import (
	"context"
	"fmt"

	"github.com/v-lozhkin/GB_Backend1_CW_GO/internal/models"
)

func (u usecase) Create(ctx context.Context, link *models.Link) error {
	if err := link.Validate(); err != nil {
		return fmt.Errorf("link's validate failed: %w", err)
	}

	if err := u.repo.Create(ctx, link); err != nil {
		return fmt.Errorf("failed to create link in repo: %w", err)
	}

	return nil
}
