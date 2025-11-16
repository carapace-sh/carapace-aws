package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_disassociateWebAclCmd = &cobra.Command{
	Use:   "disassociate-web-acl",
	Short: "This is **AWS WAF Classic Regional** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_disassociateWebAclCmd).Standalone()

	wafRegional_disassociateWebAclCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the resource from which the web ACL is being removed, either an application load balancer or Amazon API Gateway stage.")
	wafRegional_disassociateWebAclCmd.MarkFlagRequired("resource-arn")
	wafRegionalCmd.AddCommand(wafRegional_disassociateWebAclCmd)
}
