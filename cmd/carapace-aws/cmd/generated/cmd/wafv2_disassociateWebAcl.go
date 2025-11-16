package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_disassociateWebAclCmd = &cobra.Command{
	Use:   "disassociate-web-acl",
	Short: "Disassociates the specified resource from its web ACL association, if it has one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_disassociateWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_disassociateWebAclCmd).Standalone()

		wafv2_disassociateWebAclCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to disassociate from the web ACL.")
		wafv2_disassociateWebAclCmd.MarkFlagRequired("resource-arn")
	})
	wafv2Cmd.AddCommand(wafv2_disassociateWebAclCmd)
}
