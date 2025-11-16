package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_disassociateDistributionWebAcl2020_05_31Cmd = &cobra.Command{
	Use:   "disassociate-distribution-web-acl2020_05_31",
	Short: "Disassociates a distribution from the WAF web ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_disassociateDistributionWebAcl2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_disassociateDistributionWebAcl2020_05_31Cmd).Standalone()

		cloudfront_disassociateDistributionWebAcl2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution.")
		cloudfront_disassociateDistributionWebAcl2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the distribution that you're disassociating from the WAF web ACL.")
		cloudfront_disassociateDistributionWebAcl2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_disassociateDistributionWebAcl2020_05_31Cmd)
}
