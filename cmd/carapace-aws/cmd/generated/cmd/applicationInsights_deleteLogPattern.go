package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_deleteLogPatternCmd = &cobra.Command{
	Use:   "delete-log-pattern",
	Short: "Removes the specified log pattern from a `LogPatternSet`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_deleteLogPatternCmd).Standalone()

	applicationInsights_deleteLogPatternCmd.Flags().String("pattern-name", "", "The name of the log pattern.")
	applicationInsights_deleteLogPatternCmd.Flags().String("pattern-set-name", "", "The name of the log pattern set.")
	applicationInsights_deleteLogPatternCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_deleteLogPatternCmd.MarkFlagRequired("pattern-name")
	applicationInsights_deleteLogPatternCmd.MarkFlagRequired("pattern-set-name")
	applicationInsights_deleteLogPatternCmd.MarkFlagRequired("resource-group-name")
	applicationInsightsCmd.AddCommand(applicationInsights_deleteLogPatternCmd)
}
