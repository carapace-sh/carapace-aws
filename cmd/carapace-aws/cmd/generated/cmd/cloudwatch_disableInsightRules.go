package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_disableInsightRulesCmd = &cobra.Command{
	Use:   "disable-insight-rules",
	Short: "Disables the specified Contributor Insights rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_disableInsightRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_disableInsightRulesCmd).Standalone()

		cloudwatch_disableInsightRulesCmd.Flags().String("rule-names", "", "An array of the rule names to disable.")
		cloudwatch_disableInsightRulesCmd.MarkFlagRequired("rule-names")
	})
	cloudwatchCmd.AddCommand(cloudwatch_disableInsightRulesCmd)
}
