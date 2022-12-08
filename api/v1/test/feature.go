package test

import (
	"github.com/mezmerizxd/go-app/pkg/feature"
)

type TestData struct {
	Message string `json:"message"`
}

type testFeature struct {
	*feature.Feature
}

func New(parent *feature.Feature) {
	f := &testFeature{Feature: parent}

	f.Post("/test", test)
}
