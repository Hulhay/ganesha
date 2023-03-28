package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/hulhay/ganesha/graph/generated"
	"github.com/hulhay/ganesha/lib/config"
	"github.com/hulhay/ganesha/resolver"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load("./lib/config/.env")
	if err != nil {
		panic(".env is not loaded properly")
	}

	cfg := config.NewConfig()
	port := strconv.Itoa(cfg.Port())

	router := chi.NewRouter()
	router.Use(cors.Default().Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	router.Handle("/play", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/play for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
