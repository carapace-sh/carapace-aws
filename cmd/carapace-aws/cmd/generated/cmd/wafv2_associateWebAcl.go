package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_associateWebAclCmd = &cobra.Command{
	Use:   "associate-web-acl",
	Short: "Associates a web ACL with a resource, to protect the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_associateWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_associateWebAclCmd).Standalone()

		wafv2_associateWebAclCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to associate with the web ACL.")
		wafv2_associateWebAclCmd.Flags().String("web-aclarn", "", "The Amazon Resource Name (ARN) of the web ACL that you want to associate with the resource.")
		wafv2_associateWebAclCmd.MarkFlagRequired("resource-arn")
		wafv2_associateWebAclCmd.MarkFlagRequired("web-aclarn")
	})
	wafv2Cmd.AddCommand(wafv2_associateWebAclCmd)
}
