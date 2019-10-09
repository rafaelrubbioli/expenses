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
)

func Handle() http.HandlerFunc {
	return handler.GraphQL(
		graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver.Resolver{}}),

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
