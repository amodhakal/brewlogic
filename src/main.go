package main

import (
	"amodhakal/brewlogic/src/repo"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	psqlUser := os.Getenv("PSQL_USER")
	psqlPassword := os.Getenv("PSQL_PASSWORD")
	psqlHost := os.Getenv("PSQL_HOST")
	psqlPort := os.Getenv("PSQL_PORT")
	psqlDb := os.Getenv("PSQL_DB")
	psqlString := "postgres://" + psqlUser + ":" + psqlPassword + "@" + psqlHost + ":" + psqlPort + "/" + psqlDb

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, psqlString)
	if err != nil {
		log.Fatal(err)
	}

	q := repo.New(pool)
	result, err := q.GetUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
