package form

import (
	"github.com/mezmerizxd/go-app/pkg/feature"
)

type Form struct {
	Company string `json:"company"`
	Contact string `json:"contact"`
	Country string `json:"country"`
}

type formFeature struct {
	*feature.Feature
}

var table = []Form{}

func New(parent *feature.Feature) {
	f := &formFeature{Feature: parent}

	f.Post("/form/handle", handleForm)
}
