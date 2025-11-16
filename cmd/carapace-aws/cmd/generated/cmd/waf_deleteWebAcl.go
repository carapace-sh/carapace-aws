package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteWebAclCmd = &cobra.Command{
	Use:   "delete-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteWebAclCmd).Standalone()

		waf_deleteWebAclCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteWebAclCmd.Flags().String("web-aclid", "", "The `WebACLId` of the [WebACL]() that you want to delete.")
		waf_deleteWebAclCmd.MarkFlagRequired("change-token")
		waf_deleteWebAclCmd.MarkFlagRequired("web-aclid")
	})
	wafCmd.AddCommand(waf_deleteWebAclCmd)
}
