package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "create-invalidation-for-distribution-tenant2020_05_31",
	Short: "Creates an invalidation for a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd).Standalone()

	cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution tenant.")
	cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd.Flags().String("invalidation-batch", "", "")
	cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd.MarkFlagRequired("invalidation-batch")
	cloudfrontCmd.AddCommand(cloudfront_createInvalidationForDistributionTenant2020_05_31Cmd)
}
