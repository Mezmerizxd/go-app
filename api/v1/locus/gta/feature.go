package gta

import (
	"github.com/mezmerizxd/go-app/pkg/feature"
)

type gtaFeature struct {
	*feature.Feature
}

func New(parent *feature.Feature) {
	f := &gtaFeature{Feature: parent}

	f.Get("/locus/gta/{catagory}", f.getCatagory)
}
