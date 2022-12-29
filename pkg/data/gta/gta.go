package gta

import (
	"context"
	"errors"

	"github.com/mezmerizxd/go-app/pkg/cache"
)

type Masks struct {}

type Gta interface {
	GetMasks(ctx context.Context) (*[]Masks, error)
}

type gta struct {
	cache     cache.Cache
}

var ErrNoDataFound = errors.New("gta: no data found")

func New(cacheImpl cache.Cache) Gta {
	return &gta{
		cache: cacheImpl,
	}
}

func (g *gta) GetMasks(ctx context.Context) (*[]Masks, error) {
	return nil, ErrNoDataFound
}