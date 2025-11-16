package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd = &cobra.Command{
	Use:   "associate-distribution-tenant-web-acl2020_05_31",
	Short: "Associates the WAF web ACL with a distribution tenant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd).Standalone()

		cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution tenant.")
		cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd.Flags().String("if-match", "", "The current `ETag` of the distribution tenant.")
		cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd.Flags().String("web-aclarn", "", "The Amazon Resource Name (ARN) of the WAF web ACL to associate.")
		cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd.MarkFlagRequired("web-aclarn")
	})
	cloudfrontCmd.AddCommand(cloudfront_associateDistributionTenantWebAcl2020_05_31Cmd)
}
