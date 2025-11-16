package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_associateWebAclCmd = &cobra.Command{
	Use:   "associate-web-acl",
	Short: "This is **AWS WAF Classic Regional** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_associateWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_associateWebAclCmd).Standalone()

		wafRegional_associateWebAclCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the resource to be protected, either an application load balancer or Amazon API Gateway stage.")
		wafRegional_associateWebAclCmd.Flags().String("web-aclid", "", "A unique identifier (ID) for the web ACL.")
		wafRegional_associateWebAclCmd.MarkFlagRequired("resource-arn")
		wafRegional_associateWebAclCmd.MarkFlagRequired("web-aclid")
	})
	wafRegionalCmd.AddCommand(wafRegional_associateWebAclCmd)
}
