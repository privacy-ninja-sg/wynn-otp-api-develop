package database

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
	"wynn-otp-api/ent"
)

func migrateHook(next schema.Creator) schema.Creator {
	return schema.CreateFunc(func(ctx context.Context, table ...*schema.Table) error {
		return next.Create(ctx, table...)
	})
}

func CreateEntClient() *ent.Client {
	// postgres init config DSN
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL_MODE")
	pgDsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPwd, dbSSL)

	// setup postgres client connection
	pgClient, err := ent.Open("postgres", pgDsn)
	if err != nil {
		log.Fatalln(err)
	}
	//defer pgClient.Close()

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// migrate db schema
	err = pgClient.Schema.Create(ctx,
		schema.WithHooks(migrateHook),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return pgClient
}
