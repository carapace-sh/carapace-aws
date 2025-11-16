package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getWebAclCmd = &cobra.Command{
	Use:   "get-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getWebAclCmd).Standalone()

		wafRegional_getWebAclCmd.Flags().String("web-aclid", "", "The `WebACLId` of the [WebACL]() that you want to get.")
		wafRegional_getWebAclCmd.MarkFlagRequired("web-aclid")
	})
	wafRegionalCmd.AddCommand(wafRegional_getWebAclCmd)
}
