package gta

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mezmerizxd/go-app/pkg/data/gta"
	"github.com/mezmerizxd/go-app/pkg/responder"
)

type EditCatagory struct {
	Item 	string `json:"item"`
	ItemID   string    `json:"item_id"`
	TextureID string `json:"texture_id"`
}

func (f *gtaFeature) getCatagory(ctx context.Context, w responder.Responder, r *http.Request) {
	urlparam := chi.URLParam(r, "catagory")

	if urlparam == "" {
		w.Error("json", "Please enter a catagory")
		return
	}

	switch urlparam {
	case "masks":
		data, err := f.Data.Gta.GetMasks(ctx)
		if err != nil {
			w.Error("json", err.Error())
			return
		}
		w.Success("json", data)
		return

	default:
		w.Error("json", "Invalid catagory")
		return
	}
}

func (f *gtaFeature) editCatagory(ctx context.Context, w responder.Responder, r *http.Request) {
	catagory := chi.URLParam(r, "catagory")
	classID, _ := strconv.Atoi(chi.URLParam(r, "class_id"))
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if catagory == "" {
		w.Error("json", "Please enter a catagory")
		return
	}
	if classID < 0 {
		w.Error("json", "Please enter a class id")
		return
	}
	if id < 0 {
		w.Error("json", "Please enter a id")
		return
	}

	isValid := isValidCatagory(catagory)
	if !isValid {
		w.Error("json", "Invalid catagory")
		return
	}

	body := EditCatagory{}
	json.NewDecoder(r.Body).Decode(&body)

	err := f.Data.Gta.EditRowInTable(ctx, gta.SearchRowInTable{
		Catagory: catagory,
		ClassID: classID,
		ID:        id,
	}, &gta.NewDataInRow{
		Item:      body.Item,
		ItemID:    body.ItemID,
		TextureID: body.TextureID,
	})

	if err != nil {
		w.Error("json", err.Error())
		return
	}

	if err != nil {
		w.Error("json", err.Error())
		return
	}
	w.Success("json", nil)
}