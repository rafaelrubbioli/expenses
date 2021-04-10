package graphql

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"

	graphql "github.com/rafaelrubbioli/espenses/pkg/graphql/internal"
	"github.com/rafaelrubbioli/espenses/pkg/graphql/internal/resolver"
	"github.com/rafaelrubbioli/espenses/pkg/service"
)

func Handle(services service.All) http.HandlerFunc {
	return handler.GraphQL(
		graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver.New(services)}),

		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			log.Print(err)
			debug.PrintStack()
			return errors.New("internal server error")
		}),
	)
}

func HandlePlayground() http.HandlerFunc {
	return handler.Playground("Expenses", "/query")
}
