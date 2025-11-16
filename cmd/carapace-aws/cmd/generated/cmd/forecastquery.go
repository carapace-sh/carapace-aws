package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecastqueryCmd = &cobra.Command{
	Use:   "forecastquery",
	Short: "Provides APIs for creating and managing Amazon Forecast resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecastqueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecastqueryCmd).Standalone()

	})
	rootCmd.AddCommand(forecastqueryCmd)
}
