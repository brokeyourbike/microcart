package commands

import (
	"fmt"

	"github.com/brokeyourbike/microcart/business/models"
	"github.com/brokeyourbike/microcart/foundation/config"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate models with GORM",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("cannot load configuration: %w", err)
		}

		orm, err := gorm.Open(sqlite.Open(cfg.SQLite.Path), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("cannot open mysql session: %w", err)
		}

		orm.AutoMigrate(&models.Shop{})
		orm.AutoMigrate(&models.Member{})
		orm.AutoMigrate(&models.Product{})
		orm.AutoMigrate(&models.Variant{})

		fmt.Println("Models migrated!")

		return nil
	},
}
