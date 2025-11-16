package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_listManagedInsightRulesCmd = &cobra.Command{
	Use:   "list-managed-insight-rules",
	Short: "Returns a list that contains the number of managed Contributor Insights rules in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_listManagedInsightRulesCmd).Standalone()

	cloudwatch_listManagedInsightRulesCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	cloudwatch_listManagedInsightRulesCmd.Flags().String("next-token", "", "Include this value to get the next set of rules if the value was returned by the previous operation.")
	cloudwatch_listManagedInsightRulesCmd.Flags().String("resource-arn", "", "The ARN of an Amazon Web Services resource that has managed Contributor Insights rules.")
	cloudwatch_listManagedInsightRulesCmd.MarkFlagRequired("resource-arn")
	cloudwatchCmd.AddCommand(cloudwatch_listManagedInsightRulesCmd)
}
