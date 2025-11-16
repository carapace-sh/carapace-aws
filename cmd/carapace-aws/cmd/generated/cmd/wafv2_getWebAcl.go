package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getWebAclCmd = &cobra.Command{
	Use:   "get-web-acl",
	Short: "Retrieves the specified [WebACL]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_getWebAclCmd).Standalone()

		wafv2_getWebAclCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the web ACL that you want to retrieve.")
		wafv2_getWebAclCmd.Flags().String("id", "", "The unique identifier for the web ACL.")
		wafv2_getWebAclCmd.Flags().String("name", "", "The name of the web ACL.")
		wafv2_getWebAclCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	})
	wafv2Cmd.AddCommand(wafv2_getWebAclCmd)
}
