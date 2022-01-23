package context

import (
	"context"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/config"
)

type contextCfgKeys struct{}

func SetConfig(parent context.Context, cfg config.Config) context.Context {
	return context.WithValue(parent, contextCfgKeys{}, cfg)
}

func GetConfig(ctx context.Context) config.Config {
	cfg, ok := ctx.Value(contextCfgKeys{}).(config.Config)
	if !ok {
		panic("can't get config from context")
	}

	return cfg
}
