package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "create-distribution-tenant2020_05_31",
	Short: "Creates a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createDistributionTenant2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createDistributionTenant2020_05_31Cmd).Standalone()

		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("connection-group-id", "", "The ID of the connection group to associate with the distribution tenant.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("customizations", "", "Customizations for the distribution tenant.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("distribution-id", "", "The ID of the multi-tenant distribution to use for creating the distribution tenant.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("domains", "", "The domains associated with the distribution tenant.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("enabled", "", "Indicates whether the distribution tenant should be enabled when created.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("managed-certificate-request", "", "The configuration for the CloudFront managed ACM certificate request.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("name", "", "The name of the distribution tenant.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("parameters", "", "A list of parameter values to add to the resource.")
		cloudfront_createDistributionTenant2020_05_31Cmd.Flags().String("tags", "", "")
		cloudfront_createDistributionTenant2020_05_31Cmd.MarkFlagRequired("distribution-id")
		cloudfront_createDistributionTenant2020_05_31Cmd.MarkFlagRequired("domains")
		cloudfront_createDistributionTenant2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_createDistributionTenant2020_05_31Cmd)
}
