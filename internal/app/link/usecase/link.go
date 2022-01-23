package usecase

import "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link"

type usecase struct {
	repo link.Repository
}

func New(repo link.Repository) link.Usecase {
	return usecase{repo}
}
