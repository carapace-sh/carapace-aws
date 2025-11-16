package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_enableInsightRulesCmd = &cobra.Command{
	Use:   "enable-insight-rules",
	Short: "Enables the specified Contributor Insights rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_enableInsightRulesCmd).Standalone()

	cloudwatch_enableInsightRulesCmd.Flags().String("rule-names", "", "An array of the rule names to enable.")
	cloudwatch_enableInsightRulesCmd.MarkFlagRequired("rule-names")
	cloudwatchCmd.AddCommand(cloudwatch_enableInsightRulesCmd)
}
