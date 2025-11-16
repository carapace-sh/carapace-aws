package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listResourcesForWebAclCmd = &cobra.Command{
	Use:   "list-resources-for-web-acl",
	Short: "Retrieves an array of the Amazon Resource Names (ARNs) for the resources that are associated with the specified web ACL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listResourcesForWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_listResourcesForWebAclCmd).Standalone()

		wafv2_listResourcesForWebAclCmd.Flags().String("resource-type", "", "Retrieves the web ACLs that are used by the specified resource type.")
		wafv2_listResourcesForWebAclCmd.Flags().String("web-aclarn", "", "The Amazon Resource Name (ARN) of the web ACL.")
		wafv2_listResourcesForWebAclCmd.MarkFlagRequired("web-aclarn")
	})
	wafv2Cmd.AddCommand(wafv2_listResourcesForWebAclCmd)
}
