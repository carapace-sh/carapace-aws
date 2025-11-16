package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getDistributionTenant2020_05_31Cmd = &cobra.Command{
	Use:   "get-distribution-tenant2020_05_31",
	Short: "Gets information about a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getDistributionTenant2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getDistributionTenant2020_05_31Cmd).Standalone()

		cloudfront_getDistributionTenant2020_05_31Cmd.Flags().String("identifier", "", "The identifier of the distribution tenant.")
		cloudfront_getDistributionTenant2020_05_31Cmd.MarkFlagRequired("identifier")
	})
	cloudfrontCmd.AddCommand(cloudfront_getDistributionTenant2020_05_31Cmd)
}
