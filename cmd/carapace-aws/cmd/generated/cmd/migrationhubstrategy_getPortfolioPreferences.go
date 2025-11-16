package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getPortfolioPreferencesCmd = &cobra.Command{
	Use:   "get-portfolio-preferences",
	Short: "Retrieves your migration and modernization preferences.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getPortfolioPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getPortfolioPreferencesCmd).Standalone()

	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getPortfolioPreferencesCmd)
}
