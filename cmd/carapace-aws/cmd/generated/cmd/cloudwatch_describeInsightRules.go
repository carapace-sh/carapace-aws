package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_describeInsightRulesCmd = &cobra.Command{
	Use:   "describe-insight-rules",
	Short: "Returns a list of all the Contributor Insights rules in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_describeInsightRulesCmd).Standalone()

	cloudwatch_describeInsightRulesCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	cloudwatch_describeInsightRulesCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of rules.")
	cloudwatchCmd.AddCommand(cloudwatch_describeInsightRulesCmd)
}
