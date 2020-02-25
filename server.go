package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/victorneuret/graphql-go/ORM"
	"github.com/victorneuret/graphql-go/graph"
	"github.com/victorneuret/graphql-go/graph/generated"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {
	ORM.Init()
	defer ORM.DB.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
