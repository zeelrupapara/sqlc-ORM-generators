package cli

import (
	"github.com/Improwised/golang-api/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Init app initialization
func Init(cfg config.AppConfig, logger *zap.Logger) error {
	migrationCmd := GetMigrationCommandDef(cfg)

	rootCmd := &cobra.Command{Use: ""}
	rootCmd.AddCommand(&migrationCmd)
	return rootCmd.Execute()
}
