package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cjdenio/dockerenv/graph"
	"github.com/cjdenio/dockerenv/graph/generated"
	"github.com/cjdenio/dockerenv/pkg/images"
)

const defaultPort = "3000"

func main() {
	err := images.LoadImages("images/")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", http.RedirectHandler("/graphql", http.StatusTemporaryRedirect))

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			playground.Handler("GraphQL playground", "/graphql")(w, r)
		} else {
			srv.ServeHTTP(w, r)
		}
	})

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
