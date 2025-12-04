package config

import (
	"amodhakal/brewlogic/src/repo"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	PSQLUser     string
	PSQLPassword string
	PSQLHost     string
	PSQLPort     string
	PSQLDb       string
	PSQLString   string
	Query        *repo.Queries
	DBContext    context.Context
)

func init() {
	godotenv.Load()
	PSQLUser = os.Getenv("PSQL_USER")
	PSQLPassword = os.Getenv("PSQL_PASSWORD")
	PSQLHost = os.Getenv("PSQL_HOST")
	PSQLPort = os.Getenv("PSQL_PORT")
	PSQLDb = os.Getenv("PSQL_DB")
	PSQLString = "postgres://" + PSQLUser + ":" + PSQLPassword + "@" + PSQLHost + ":" + PSQLPort + "/" + PSQLDb

	DBContext = context.Background()
	pool, err := pgxpool.New(DBContext, PSQLString)
	if err != nil {
		log.Fatal(err)
	}

	Query = repo.New(pool)
}
