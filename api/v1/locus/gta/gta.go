package gta

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mezmerizxd/go-app/pkg/responder"
)

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