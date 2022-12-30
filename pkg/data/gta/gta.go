package gta

import (
	"context"
	"errors"
	"log"

	"github.com/mezmerizxd/go-app/pkg/cache"
	"github.com/mezmerizxd/go-app/pkg/firebase"
)

type MasksItem struct {
	ID		  int `json:"id"`
	Item 	string `json:"item"`
	ItemClass string `json:"item_class"`
	ItemID   string    `json:"item_id"`
	TextureID string `json:"texture_id"`
}

type Masks struct {
	ClassItems []MasksItem `json:"class_items"`
	ClassName  string  `json:"class_name"`
}

type Gta interface {
	GetMasks(ctx context.Context) (*[]Masks, error)
}

type gta struct {
	cache     cache.Cache
	firebase firebase.Firebase
}

var ErrNoDataFound = errors.New("gta: no data found")

func New(cacheImpl cache.Cache, firebaseImpl firebase.Firebase) Gta {
	return &gta{
		cache: cacheImpl,
		firebase: firebaseImpl,
	}
}

func (g *gta) GetMasks(ctx context.Context) (*[]Masks, error) {
	var masks []Masks

	client := g.firebase.Database()
	ref := client.NewRef("locus/masks")

	if err := ref.Get(ctx, &masks); err != nil {
        log.Println("Error reading value:", err)
		return nil, ErrNoDataFound
	}

	return &masks, nil
}