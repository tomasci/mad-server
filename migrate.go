package main

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// connecting to database
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresDb := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	dsn := url.URL{
		Scheme: "postgres",
		Host:   postgresHost,
		Path:   postgresDb,
		User:   url.UserPassword(postgresUser, postgresPassword),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Printf("failed to connect database: %v", err)
	}

	// migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// todo: migrations of different entities must be split to different files
		{
			ID: "create_users_table",
			Migrate: func(tx *gorm.DB) error {
				type users struct {
					ID        uuid.UUID `gorm:"type:uuid;primaryKey;uniqueIndex" json:"id"`
					Username  string    `gorm:"unique;not null" json:"username"`
					Password  string    `json:"password"`
					Email     string    `gorm:"unique;not null" json:"email"`
					CreatedAt time.Time `json:"created_at"`
					UpdatedAt time.Time `json:"updated_at"`
					DeletedAt time.Time `gorm:"index" json:"deleted_at"`
				}

				return tx.Migrator().CreateTable(&users{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
		//        {
		//            ID: "users_rename_Name_to_Username",
		//            Migrate: func(tx *gorm.DB) error {
		//                return tx.Migrator().RenameColumn("users", "name", "username")
		//            },
		//            Rollback: func(tx *gorm.DB) error {
		//                return tx.Migrator().RenameColumn("users", "username", "name")
		//            },
		//        },
		{
			ID: "create_todos_table",
			Migrate: func(tx *gorm.DB) error {
				type todos struct {
					ID        uuid.UUID `gorm:"type:uuid;primaryKey;uniqueIndex" json:"id"`
					CreatedAt time.Time `json:"created_at"`
					UpdatedAt time.Time `json:"updated_at"`
					DeletedAt time.Time `gorm:"index" json:"deleted_at"`
				}

				return tx.Migrator().CreateTable(&todos{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("todos")
			},
		},
		{
			ID: "create_users_todos_table",
			Migrate: func(tx *gorm.DB) error {
				type usersTodos struct {
					UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
					TodoID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"todo_id"`
					CreatedAt time.Time `json:"created_at"`
				}
				return tx.Migrator().CreateTable(&usersTodos{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users_todos")
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v\n", err)
	}
	fmt.Println("Migrations: ok.")
}
