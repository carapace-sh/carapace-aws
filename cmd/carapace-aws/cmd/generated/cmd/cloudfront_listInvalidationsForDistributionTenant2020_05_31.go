package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "list-invalidations-for-distribution-tenant2020_05_31",
	Short: "Lists the invalidations for a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd).Standalone()

		cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution tenant.")
		cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd.Flags().String("marker", "", "Use this parameter when paginating results to indicate where to begin in your list of invalidation batches.")
		cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of invalidations to return for the distribution tenant.")
		cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_listInvalidationsForDistributionTenant2020_05_31Cmd)
}
