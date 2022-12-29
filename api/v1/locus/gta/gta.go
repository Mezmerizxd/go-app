package gta

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/go-chi/chi/v5"
	"github.com/mezmerizxd/go-app/pkg/responder"
)

type Post struct {
	Authorization string `json:"authorization,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Email  string `json:"email,omitempty"`
}

func (f *gtaFeature) getCatagory(ctx context.Context, w responder.Responder, r *http.Request) {
	urlparam := chi.URLParam(r, "catagory")

	if urlparam == "" {
		w.Error("json", "Please enter a catagory")
		return
	}

	data := make(map[string]interface{})

	conf := &firebase.Config{ProjectID: "mezmerizxd-social-app"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
    	log.Fatalln("Error initializing database client:", err)
	}
	ref := client.NewRef("social_app_v2/user_data/452802201068579")
	var post Post
	if err := ref.Get(ctx, &post); err != nil {
        log.Fatalln("Error reading value:", err)
	}
	data["masks"] = post

	// client, err := app.Firestore(ctx)
	// if err != nil {
	// log.Fatalln(err)
	// }
	// defer client.Close()

	// dsnap, err := client.Collection("test").Doc("abc").Get(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// m := dsnap.Data()
	// data["masks"] = m


	// switch urlparam {
	// case "masks":
	// 	data, err := f.Data.Gta.GetMasks(ctx)
	// 	if err != nil {
	// 		w.Error("json", "No masks found")
	// 		return
	// 	}
	// 	w.Success("json", data)
	// 	return

	// default:
	// 	w.Error("json", "Invalid catagory")
	// 	return
	// }
	w.Success("json", data)
}