package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeContributorInsightsCmd = &cobra.Command{
	Use:   "describe-contributor-insights",
	Short: "Returns information about contributor insights for a given table or global secondary index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeContributorInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_describeContributorInsightsCmd).Standalone()

		dynamodb_describeContributorInsightsCmd.Flags().String("index-name", "", "The name of the global secondary index to describe, if applicable.")
		dynamodb_describeContributorInsightsCmd.Flags().String("table-name", "", "The name of the table to describe.")
		dynamodb_describeContributorInsightsCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_describeContributorInsightsCmd)
}
