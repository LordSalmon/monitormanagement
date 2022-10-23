package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
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
	fmt.Println("Initializing database...")
	ConnectToDatabase()
	CreateWindowTable()
}

func ConnectToDatabase() {
	fmt.Println("Connecting to database...")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	Db = bun.NewDB(sqldb, pgdialect.New())
	fmt.Println("Connected to database successfully!")
}

func CreateWindowTable() {
	fmt.Println("Creating window table...")
	ctx := context.Background()
	_, err := Db.NewCreateTable().Model((*Window)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Window table created successfully!")
}

func InsertWindow(window Window) {
	ctx := context.Background()
	_, err := Db.NewInsert().Model(&window).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
