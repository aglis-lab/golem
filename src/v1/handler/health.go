package handler

import (
	"net/http"

	"github.com/aglis-lab/golem/src/app"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := app.Tracer().Start(r.Context(), "HealthHandler")
		defer span.End()

		w.Write([]byte("ok"))
	}
}
