package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/oskanberg/reqcheck"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	http.Handle("/playground", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(
		reqcheck.NewExecutableSchema(reqcheck.Config{Resolvers: &reqcheck.Resolver{}}),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
