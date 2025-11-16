package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd = &cobra.Command{
	Use:   "disassociate-distribution-tenant-web-acl2020_05_31",
	Short: "Disassociates a distribution tenant from the WAF web ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd).Standalone()

		cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution tenant.")
		cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd.Flags().String("if-match", "", "The current version of the distribution tenant that you're disassociating from the WAF web ACL.")
		cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_disassociateDistributionTenantWebAcl2020_05_31Cmd)
}
