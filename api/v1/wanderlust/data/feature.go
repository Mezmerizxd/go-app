package data

import (
	"github.com/mezmerizxd/go-app/pkg/feature"
)

type Data struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Other string `json:"other"`
}

type dataFeature struct {
	*feature.Feature
}

func New(parent *feature.Feature) {
	f := &dataFeature{Feature: parent}

	f.Post("/wanderlust/data", handleData)
}
