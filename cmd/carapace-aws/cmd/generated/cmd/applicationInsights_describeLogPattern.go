package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeLogPatternCmd = &cobra.Command{
	Use:   "describe-log-pattern",
	Short: "Describe a specific log pattern from a `LogPatternSet`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeLogPatternCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_describeLogPatternCmd).Standalone()

		applicationInsights_describeLogPatternCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_describeLogPatternCmd.Flags().String("pattern-name", "", "The name of the log pattern.")
		applicationInsights_describeLogPatternCmd.Flags().String("pattern-set-name", "", "The name of the log pattern set.")
		applicationInsights_describeLogPatternCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_describeLogPatternCmd.MarkFlagRequired("pattern-name")
		applicationInsights_describeLogPatternCmd.MarkFlagRequired("pattern-set-name")
		applicationInsights_describeLogPatternCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_describeLogPatternCmd)
}
