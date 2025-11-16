package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_associateDistributionWebAcl2020_05_31Cmd = &cobra.Command{
	Use:   "associate-distribution-web-acl2020_05_31",
	Short: "Associates the WAF web ACL with a distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_associateDistributionWebAcl2020_05_31Cmd).Standalone()

	cloudfront_associateDistributionWebAcl2020_05_31Cmd.Flags().String("id", "", "The ID of the distribution.")
	cloudfront_associateDistributionWebAcl2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the distribution that you're associating with the WAF web ACL.")
	cloudfront_associateDistributionWebAcl2020_05_31Cmd.Flags().String("web-aclarn", "", "The Amazon Resource Name (ARN) of the WAF web ACL to associate.")
	cloudfront_associateDistributionWebAcl2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_associateDistributionWebAcl2020_05_31Cmd.MarkFlagRequired("web-aclarn")
	cloudfrontCmd.AddCommand(cloudfront_associateDistributionWebAcl2020_05_31Cmd)
}
