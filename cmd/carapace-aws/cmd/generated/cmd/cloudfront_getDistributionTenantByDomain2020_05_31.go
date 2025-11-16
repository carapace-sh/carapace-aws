package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getDistributionTenantByDomain2020_05_31Cmd = &cobra.Command{
	Use:   "get-distribution-tenant-by-domain2020_05_31",
	Short: "Gets information about a distribution tenant by the associated domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getDistributionTenantByDomain2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getDistributionTenantByDomain2020_05_31Cmd).Standalone()

		cloudfront_getDistributionTenantByDomain2020_05_31Cmd.Flags().String("domain", "", "A domain name associated with the target distribution tenant.")
		cloudfront_getDistributionTenantByDomain2020_05_31Cmd.MarkFlagRequired("domain")
	})
	cloudfrontCmd.AddCommand(cloudfront_getDistributionTenantByDomain2020_05_31Cmd)
}
