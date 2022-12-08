package count

import (
	"github.com/mezmerizxd/go-app/pkg/feature"
)

type Count struct {
	Count int `json:"count"`
}

type countFeature struct {
	*feature.Feature
}

var count = 0

func New(parent *feature.Feature) {
	f := &countFeature{Feature: parent}

	f.Post("/handle_count", handleCount)
	f.Get("/fetch_count", fetchCount)
}
