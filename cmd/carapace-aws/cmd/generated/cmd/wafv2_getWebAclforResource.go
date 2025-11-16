package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getWebAclforResourceCmd = &cobra.Command{
	Use:   "get-web-aclfor-resource",
	Short: "Retrieves the [WebACL]() for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getWebAclforResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_getWebAclforResourceCmd).Standalone()

		wafv2_getWebAclforResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose web ACL you want to retrieve.")
		wafv2_getWebAclforResourceCmd.MarkFlagRequired("resource-arn")
	})
	wafv2Cmd.AddCommand(wafv2_getWebAclforResourceCmd)
}
