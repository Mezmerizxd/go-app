package gta

import (
	"context"

	"github.com/mezmerizxd/go-app/pkg/feature"
	"github.com/pkgz/websocket"
)

type gtaFeature struct {
	*feature.Feature
}

func New(parent *feature.Feature) {
	f := &gtaFeature{Feature: parent}

	f.Get("/locus/gta/{catagory}", f.getCatagory)
	f.Post("/locus/gta/{catagory}/edit/{class_id}/{id}", f.editCatagory)
}

func isValidCatagory(catagory string) bool {
	switch catagory {
	case "masks":
		return true
	case "dev_masks":
		return true
	default:
		return false
	}
}

func testSocket(ctx context.Context, conn *websocket.Conn, msg websocket.Message) {

}