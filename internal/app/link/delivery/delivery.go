package delivery

import "github.com/v-lozhkin/GB_Backend1_CW_GO/internal/app/link"

type delivery struct {
	links link.Usecase
}

func New(links link.Usecase) link.Delivery {
	return delivery{
		links: links,
	}
}
