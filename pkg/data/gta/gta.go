package gta

import (
	"context"
	"errors"
	"log"
	"strconv"

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
	ClassID   int     `json:"class_id"`
}

type SearchRowInTable struct {
	Catagory string `json:"catagory"`
	ClassID int `json:"class_id"`
	ID        int    `json:"id"`
}

type NewDataInRow struct {
	Item 	string `json:"item"`
	ItemID   string    `json:"item_id"`
	TextureID string `json:"texture_id"`
}

type Gta interface {
	GetMasks(ctx context.Context) (*[]Masks, error)
	EditRowInTable(ctx context.Context, search SearchRowInTable, data *NewDataInRow) (error)
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

func (g *gta) EditRowInTable(ctx context.Context, search SearchRowInTable, data *NewDataInRow) (error) {
	var masks MasksItem

	client := g.firebase.Database()
	ref := client.NewRef("locus/" + search.Catagory + "/" + strconv.Itoa(search.ClassID) + "/class_items/" + strconv.Itoa(search.ID))
	
	if err := ref.Get(ctx, &masks); err != nil {
		log.Println("Error reading value:", err)
		return ErrNoDataFound
	}

	// Edit data
	if data.Item != "" {
		masks.Item = data.Item
	}
	if data.ItemID != "" {
		masks.ItemID = data.ItemID
	}
	if data.TextureID != "" {
		masks.TextureID = data.TextureID
	}

	if err := ref.Set(ctx, masks); err != nil {
		log.Println("Error setting value:", err)
		return ErrNoDataFound
	}

	return nil
}