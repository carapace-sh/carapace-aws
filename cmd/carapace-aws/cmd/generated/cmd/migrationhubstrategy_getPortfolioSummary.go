package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getPortfolioSummaryCmd = &cobra.Command{
	Use:   "get-portfolio-summary",
	Short: "Retrieves overall summary including the number of servers to rehost and the overall number of anti-patterns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getPortfolioSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getPortfolioSummaryCmd).Standalone()

	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getPortfolioSummaryCmd)
}
