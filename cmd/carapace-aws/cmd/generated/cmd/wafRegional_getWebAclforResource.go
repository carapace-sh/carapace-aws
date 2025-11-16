package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getWebAclforResourceCmd = &cobra.Command{
	Use:   "get-web-aclfor-resource",
	Short: "This is **AWS WAF Classic Regional** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getWebAclforResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getWebAclforResourceCmd).Standalone()

		wafRegional_getWebAclforResourceCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the resource for which to get the web ACL, either an application load balancer or Amazon API Gateway stage.")
		wafRegional_getWebAclforResourceCmd.MarkFlagRequired("resource-arn")
	})
	wafRegionalCmd.AddCommand(wafRegional_getWebAclforResourceCmd)
}
