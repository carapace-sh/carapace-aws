package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listResourcesForWebAclCmd = &cobra.Command{
	Use:   "list-resources-for-web-acl",
	Short: "This is **AWS WAF Classic Regional** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listResourcesForWebAclCmd).Standalone()

	wafRegional_listResourcesForWebAclCmd.Flags().String("resource-type", "", "The type of resource to list, either an application load balancer or Amazon API Gateway.")
	wafRegional_listResourcesForWebAclCmd.Flags().String("web-aclid", "", "The unique identifier (ID) of the web ACL for which to list the associated resources.")
	wafRegional_listResourcesForWebAclCmd.MarkFlagRequired("web-aclid")
	wafRegionalCmd.AddCommand(wafRegional_listResourcesForWebAclCmd)
}
