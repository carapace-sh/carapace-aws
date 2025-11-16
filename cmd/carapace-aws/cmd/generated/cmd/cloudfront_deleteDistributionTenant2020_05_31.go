package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "delete-distribution-tenant2020_05_31",
	Short: "Deletes a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteDistributionTenant2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deleteDistributionTenant2020_05_31Cmd).Standalone()

		cloudfront_deleteDistributionTenant2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution tenant to delete.")
		cloudfront_deleteDistributionTenant2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the distribution tenant.")
		cloudfront_deleteDistributionTenant2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_deleteDistributionTenant2020_05_31Cmd.MarkFlagRequired("if-match")
	})
	cloudfrontCmd.AddCommand(cloudfront_deleteDistributionTenant2020_05_31Cmd)
}
