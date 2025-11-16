package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_updateLogPatternCmd = &cobra.Command{
	Use:   "update-log-pattern",
	Short: "Adds a log pattern to a `LogPatternSet`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_updateLogPatternCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_updateLogPatternCmd).Standalone()

		applicationInsights_updateLogPatternCmd.Flags().String("pattern", "", "The log pattern.")
		applicationInsights_updateLogPatternCmd.Flags().String("pattern-name", "", "The name of the log pattern.")
		applicationInsights_updateLogPatternCmd.Flags().String("pattern-set-name", "", "The name of the log pattern set.")
		applicationInsights_updateLogPatternCmd.Flags().String("rank", "", "Rank of the log pattern.")
		applicationInsights_updateLogPatternCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_updateLogPatternCmd.MarkFlagRequired("pattern-name")
		applicationInsights_updateLogPatternCmd.MarkFlagRequired("pattern-set-name")
		applicationInsights_updateLogPatternCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_updateLogPatternCmd)
}
