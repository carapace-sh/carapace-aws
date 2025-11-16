package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "get-invalidation-for-distribution-tenant2020_05_31",
	Short: "Gets information about a specific invalidation for a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd).Standalone()

	cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd.Flags().String("distribution-tenant-id", "", "The ID of the distribution tenant.")
	cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd.Flags().String("id", "", "The ID of the invalidation to retrieve.")
	cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd.MarkFlagRequired("distribution-tenant-id")
	cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getInvalidationForDistributionTenant2020_05_31Cmd)
}
