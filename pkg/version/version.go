package version

import (
	"github.com/go-chi/chi/v5"
	"github.com/mezmerizxd/go-app/pkg/cache"
	"github.com/mezmerizxd/go-app/pkg/data"
	"github.com/mezmerizxd/go-app/pkg/data/gta"
	"github.com/mezmerizxd/go-app/pkg/data/users"
	"github.com/mezmerizxd/go-app/pkg/feature"
	"github.com/mezmerizxd/go-app/pkg/firebase"
	"github.com/pkgz/websocket"
)

type Version struct {
	*chi.Mux
	WsMux *websocket.Server

	featureCfg *feature.Config
	Routes     []*feature.Route
}

type Config struct {
	Cache cache.Cache
	Data  *data.Data
	Firebase firebase.Firebase
	Users users.Users
	Gta gta.Gta
}

func NewRoute(cfg *Config) *Version {
	r := chi.NewRouter()

	return &Version{
		Mux: r,
		featureCfg: &feature.Config{
			Cache: cfg.Cache,
			Data:  cfg.Data,
			Firebase: cfg.Firebase,
			Users: cfg.Users,
			Gta: cfg.Gta,
		},
	}
}

func (v *Version) RegisterRoute(factory func(*feature.Feature)) {
	f := feature.New(v.featureCfg)
	factory(f)

	for _, r := range f.Routes() {
		r.InjectRoute(v.Mux)
		v.Routes = append(v.Routes, r)
	}
}