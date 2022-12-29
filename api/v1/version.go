package v1

import (
	"github.com/mezmerizxd/go-app/api/v1/locus/gta"
	"github.com/mezmerizxd/go-app/api/v1/test"
	"github.com/mezmerizxd/go-app/api/v1/wanderlust/count"
	"github.com/mezmerizxd/go-app/api/v1/wanderlust/data"
	"github.com/mezmerizxd/go-app/api/v1/wanderlust/form"
	"github.com/mezmerizxd/go-app/api/v1/wanderlust/users"
	"github.com/mezmerizxd/go-app/pkg/version"
)

func New(cfg *version.Config) *version.Version {
	v := version.New(cfg)

	v.Register(test.New)
	v.Register(data.New)
	v.Register(count.New)
	v.Register(form.New)
	v.Register(users.New)
	v.Register(gta.New)

	return v
}
