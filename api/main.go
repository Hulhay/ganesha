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
	"github.com/hulhay/ganesha/repositories/postgresql"
	"github.com/hulhay/ganesha/resolver"
	"github.com/hulhay/ganesha/services/genre"
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

	svc := initService(cfg)
	resolver := initResolver(svc)

	router := chi.NewRouter()
	router.Use(cors.Default().Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/play", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/play for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type serviceProvider struct {
	genre *genre.GenreService
}

func initService(cfg *config.Config) serviceProvider {
	genreRepo := postgresql.NewGenreRepository(cfg)

	genreService := genre.NewGenreService(genre.GenreOptions{
		GenreRepo: genreRepo,
	})

	return serviceProvider{
		genre: genreService,
	}
}

func initResolver(svc serviceProvider) *resolver.Resolver {
	return resolver.New(resolver.Options{
		GenreService: svc.genre,
	})
}
