package db

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mokhlesurr031/goblog/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202305191200", // Unique migration ID
			Migrate: func(tx *gorm.DB) error {
				// Auto-migrate the User model
				return tx.AutoMigrate(&models.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Drop the users table
				return tx.Migrator().DropTable("users")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}
	log.Printf("Migration did run successfully")
	return nil
}
