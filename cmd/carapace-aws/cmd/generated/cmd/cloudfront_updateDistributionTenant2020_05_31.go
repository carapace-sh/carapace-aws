package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "update-distribution-tenant2020_05_31",
	Short: "Updates a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateDistributionTenant2020_05_31Cmd).Standalone()

	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("connection-group-id", "", "The ID of the target connection group.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("customizations", "", "Customizations for the distribution tenant.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("distribution-id", "", "The ID for the multi-tenant distribution.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("domains", "", "The domains to update for the distribution tenant.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("enabled", "", "Indicates whether the distribution tenant should be updated to an enabled state.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution tenant.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the distribution tenant to update.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("managed-certificate-request", "", "An object that contains the CloudFront managed ACM certificate request.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.Flags().String("parameters", "", "A list of parameter values to add to the resource.")
	cloudfront_updateDistributionTenant2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_updateDistributionTenant2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfrontCmd.AddCommand(cloudfront_updateDistributionTenant2020_05_31Cmd)
}
