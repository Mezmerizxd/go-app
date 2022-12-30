package firebase

import (
	"context"
	"log"

	FB "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/mezmerizxd/go-app/pkg/cache"
	"google.golang.org/api/option"
)

type Config struct {
	Cache cache.Cache
	DatabaseUrl string
	CredentialPath string
}

type Firebase interface {
	Database() db.Client
}

type firebase struct {
	cache cache.Cache
	app *FB.App
}

// var ErrInitializingApp = errors.New("firebase: error initializing app")
// var ErrInitializingDatabase = errors.New("firebase: error initializing database")

func New(cfg *Config) Firebase {
	conf := &FB.Config{
		DatabaseURL: cfg.DatabaseUrl,
	}
	
	opt := option.WithCredentialsFile(cfg.CredentialPath)
	app, err := FB.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &firebase{
		cache: cfg.Cache,
		app: app,
	}
}

func (f *firebase) Database() db.Client {
	client, err := f.app.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return *client
}