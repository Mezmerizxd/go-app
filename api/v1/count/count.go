package count

import (
	"context"
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/responder"
)

func handleCount(ctx context.Context, w responder.Responder, r *http.Request) {
	count = count + 1
	body := Count{
		Count: count,
	}
	w.Success("json", body)
}

func fetchCount(ctx context.Context, w responder.Responder, r *http.Request) {
	body := Count{
		Count: count,
	}
	w.Success("json", body)
}