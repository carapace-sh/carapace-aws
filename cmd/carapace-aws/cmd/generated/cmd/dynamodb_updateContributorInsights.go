package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateContributorInsightsCmd = &cobra.Command{
	Use:   "update-contributor-insights",
	Short: "Updates the status for contributor insights for a specific table or index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateContributorInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_updateContributorInsightsCmd).Standalone()

		dynamodb_updateContributorInsightsCmd.Flags().String("contributor-insights-action", "", "Represents the contributor insights action.")
		dynamodb_updateContributorInsightsCmd.Flags().String("contributor-insights-mode", "", "Specifies whether to track all access and throttled events or throttled events only for the DynamoDB table or index.")
		dynamodb_updateContributorInsightsCmd.Flags().String("index-name", "", "The global secondary index name, if applicable.")
		dynamodb_updateContributorInsightsCmd.Flags().String("table-name", "", "The name of the table.")
		dynamodb_updateContributorInsightsCmd.MarkFlagRequired("contributor-insights-action")
		dynamodb_updateContributorInsightsCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_updateContributorInsightsCmd)
}
