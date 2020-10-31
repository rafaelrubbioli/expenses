package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/example/dataloader"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rafaelrubbioli/espenses/pkg/graphql"
	"github.com/rafaelrubbioli/espenses/pkg/lib/db"
	"github.com/rafaelrubbioli/espenses/pkg/service"
)

func main() {
	router := chi.NewRouter()
	router.Use(dataloader.LoaderMiddleware)

	dbconn, err := db.NewMySQL()
	if err != nil {
		log.Fatal(err)
	}

	services := service.All{
		UserService: service.NewUserService(dbconn),
	}

	router.Handle("/", graphql.HandlePlayground())
	router.Handle("/query", graphql.Handle(services))

	log.Println("connect to http://localhost:5000/ for graphql playground")
	log.Fatal(http.ListenAndServe(":5000", router))
}
