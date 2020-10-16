package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/example/dataloader"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rafaelrubbioli/espenses/pkg/graphql"
)

func main() {
	router := chi.NewRouter()
	router.Use(dataloader.LoaderMiddleware)

	router.Handle("/", graphql.HandlePlayground())
	router.Handle("/query", graphql.Handle())

	log.Println("connect to http://localhost:5000/ for graphql playground")
	log.Fatal(http.ListenAndServe(":5000", router))
}
