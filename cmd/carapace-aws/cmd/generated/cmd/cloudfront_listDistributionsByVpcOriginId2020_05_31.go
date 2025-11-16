package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-vpc-origin-id2020_05_31",
	Short: "List CloudFront distributions by their VPC origin ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd).Standalone()

		cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd.Flags().String("marker", "", "The marker associated with the VPC origin distributions list.")
		cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of items included in the list.")
		cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd.Flags().String("vpc-origin-id", "", "The VPC origin ID.")
		cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd.MarkFlagRequired("vpc-origin-id")
	})
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByVpcOriginId2020_05_31Cmd)
}
