//go:generate weaver generate ./...

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

type app struct {
	weaver.Implements[weaver.Main]

	lis      weaver.Listener
	searcher weaver.Ref[Searcher]
}

func main() {
	ctx := context.Background()
	if err := weaver.Run(ctx, run); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, a *app) error {
	r := a.newEmojisRoutes()
	return http.Serve(a.lis, r)
}

func (a app) newEmojisRoutes() http.Handler {
	mux := http.NewServeMux()
	// a request of the type http://localhost:8080/search?q=black+cat will be done.
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		ctx := r.Context()
		results, err := a.searcher.Get().Search(ctx, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(&results)
	})
	return mux
}
