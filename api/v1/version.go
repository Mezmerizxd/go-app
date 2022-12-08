package v1

import (
	"github.com/mezmerizxd/go-app/api/v1/count"
	"github.com/mezmerizxd/go-app/api/v1/data"
	"github.com/mezmerizxd/go-app/api/v1/test"
	"github.com/mezmerizxd/go-app/pkg/version"
)

func New(cfg *version.Config) *version.Version {
	v := version.New(cfg)

	v.Register(test.New)
	v.Register(data.New)
	v.Register(count.New)

	return v
}