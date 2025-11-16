package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_createLogPatternCmd = &cobra.Command{
	Use:   "create-log-pattern",
	Short: "Adds an log pattern to a `LogPatternSet`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_createLogPatternCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_createLogPatternCmd).Standalone()

		applicationInsights_createLogPatternCmd.Flags().String("pattern", "", "The log pattern.")
		applicationInsights_createLogPatternCmd.Flags().String("pattern-name", "", "The name of the log pattern.")
		applicationInsights_createLogPatternCmd.Flags().String("pattern-set-name", "", "The name of the log pattern set.")
		applicationInsights_createLogPatternCmd.Flags().String("rank", "", "Rank of the log pattern.")
		applicationInsights_createLogPatternCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_createLogPatternCmd.MarkFlagRequired("pattern")
		applicationInsights_createLogPatternCmd.MarkFlagRequired("pattern-name")
		applicationInsights_createLogPatternCmd.MarkFlagRequired("pattern-set-name")
		applicationInsights_createLogPatternCmd.MarkFlagRequired("rank")
		applicationInsights_createLogPatternCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_createLogPatternCmd)
}
