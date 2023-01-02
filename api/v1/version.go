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
	v := version.NewRoute(cfg)

	v.RegisterRoute(test.New)
	v.RegisterRoute(data.New)
	v.RegisterRoute(count.New)
	v.RegisterRoute(form.New)
	v.RegisterRoute(users.New)
	v.RegisterRoute(gta.New)

	return v
}