package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_putPortfolioPreferencesCmd = &cobra.Command{
	Use:   "put-portfolio-preferences",
	Short: "Saves the specified migration and modernization preferences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_putPortfolioPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_putPortfolioPreferencesCmd).Standalone()

		migrationhubstrategy_putPortfolioPreferencesCmd.Flags().String("application-mode", "", "The classification for application component types.")
		migrationhubstrategy_putPortfolioPreferencesCmd.Flags().String("application-preferences", "", "The transformation preferences for non-database applications.")
		migrationhubstrategy_putPortfolioPreferencesCmd.Flags().String("database-preferences", "", "The transformation preferences for database applications.")
		migrationhubstrategy_putPortfolioPreferencesCmd.Flags().String("prioritize-business-goals", "", "The rank of the business goals based on priority.")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_putPortfolioPreferencesCmd)
}
