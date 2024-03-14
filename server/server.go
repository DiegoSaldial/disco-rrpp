package main

import (
	"auth/database/xauth"
	"auth/graph"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func conexion() *sql.DB {
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	loc := "America%2FLa_Paz"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=%s", dbuser, dbpass, dbhost, dbname, loc)
	db, err := sql.Open("mysql", dsn)
	fmt.Println(dsn)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	er := db.Ping()
	if er != nil {
		panic(er.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

const defaultPort = "8020"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := conexion()
	router := chi.NewRouter()
	router.Use(xauth.AuthMiddleware(db))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	show_playground := os.Getenv("PLAYGROUND")
	if show_playground == "1" {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
