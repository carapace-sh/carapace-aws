package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_deleteInsightRulesCmd = &cobra.Command{
	Use:   "delete-insight-rules",
	Short: "Permanently deletes the specified Contributor Insights rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_deleteInsightRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_deleteInsightRulesCmd).Standalone()

		cloudwatch_deleteInsightRulesCmd.Flags().String("rule-names", "", "An array of the rule names to delete.")
		cloudwatch_deleteInsightRulesCmd.MarkFlagRequired("rule-names")
	})
	cloudwatchCmd.AddCommand(cloudwatch_deleteInsightRulesCmd)
}
