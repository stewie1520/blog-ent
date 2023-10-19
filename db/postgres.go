package db

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/stewie1520/blog_ent/ent"

	_ "github.com/lib/pq"
	"github.com/stewie1520/blog_ent/log"
)

func NewPostgresDBX(connectionURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", connectionURL)
	if err != nil {
		log.S().Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
