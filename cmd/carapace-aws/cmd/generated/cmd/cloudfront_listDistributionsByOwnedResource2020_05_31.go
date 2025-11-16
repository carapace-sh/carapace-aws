package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByOwnedResource2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-owned-resource2020_05_31",
	Short: "Lists the CloudFront distributions that are associated with the specified resource that you own.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByOwnedResource2020_05_31Cmd).Standalone()

	cloudfront_listDistributionsByOwnedResource2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of distributions.")
	cloudfront_listDistributionsByOwnedResource2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distributions to return.")
	cloudfront_listDistributionsByOwnedResource2020_05_31Cmd.Flags().String("resource-arn", "", "The ARN of the CloudFront resource that you've shared with other Amazon Web Services accounts.")
	cloudfront_listDistributionsByOwnedResource2020_05_31Cmd.MarkFlagRequired("resource-arn")
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByOwnedResource2020_05_31Cmd)
}
