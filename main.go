package main

import (
	"log"
	"net/http"

	"github.com/rafaelrubbioli/espenses/generated"

	"github.com/rafaelrubbioli/espenses/resolver"

	"github.com/99designs/gqlgen/example/dataloader"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Use(dataloader.LoaderMiddleware)

	router.Handle("/", handler.Playground("Expenses", "/query"))
	router.Handle("/query", handler.GraphQL(
		generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}),
	))

	log.Println("connect to http://localhost:5000/ for graphql playground")
	log.Fatal(http.ListenAndServe(":5000", router))
}
