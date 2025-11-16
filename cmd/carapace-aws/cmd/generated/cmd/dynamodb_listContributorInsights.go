package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listContributorInsightsCmd = &cobra.Command{
	Use:   "list-contributor-insights",
	Short: "Returns a list of ContributorInsightsSummary for a table and all its global secondary indexes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listContributorInsightsCmd).Standalone()

	dynamodb_listContributorInsightsCmd.Flags().String("max-results", "", "Maximum number of results to return per page.")
	dynamodb_listContributorInsightsCmd.Flags().String("next-token", "", "A token to for the desired page, if there is one.")
	dynamodb_listContributorInsightsCmd.Flags().String("table-name", "", "The name of the table.")
	dynamodbCmd.AddCommand(dynamodb_listContributorInsightsCmd)
}
