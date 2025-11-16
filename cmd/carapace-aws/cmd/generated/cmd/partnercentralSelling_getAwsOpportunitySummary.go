package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getAwsOpportunitySummaryCmd = &cobra.Command{
	Use:   "get-aws-opportunity-summary",
	Short: "Retrieves a summary of an AWS Opportunity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getAwsOpportunitySummaryCmd).Standalone()

	partnercentralSelling_getAwsOpportunitySummaryCmd.Flags().String("catalog", "", "Specifies the catalog in which the AWS Opportunity is located.")
	partnercentralSelling_getAwsOpportunitySummaryCmd.Flags().String("related-opportunity-identifier", "", "The unique identifier for the related partner opportunity.")
	partnercentralSelling_getAwsOpportunitySummaryCmd.MarkFlagRequired("catalog")
	partnercentralSelling_getAwsOpportunitySummaryCmd.MarkFlagRequired("related-opportunity-identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getAwsOpportunitySummaryCmd)
}
