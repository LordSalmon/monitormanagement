package database

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var Db *bun.DB

type Window struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	WindowId  int
	Title     string
	Program   string
	X         int
	Y         int
	Width     int
	Height    int
	Depth     int
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func InitDatabase() {
	log.Info("Initializing database...")
	ConnectToDatabase()
	CreateWindowTable()
}

func ConnectToDatabase() {
	log.Info("Connecting to database...")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	Db = bun.NewDB(sqldb, pgdialect.New())
	log.Info("Connected to database successfully!")
}

func CreateWindowTable() {
	log.Info("Creating window table...")
	ctx := context.Background()
	_, err := Db.NewCreateTable().Model((*Window)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}
	log.Info("Window table created successfully!")
}

func InsertWindow(window Window) {
	ctx := context.Background()
	_, err := Db.NewInsert().Model(&window).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
