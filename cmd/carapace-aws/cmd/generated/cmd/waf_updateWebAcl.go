package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateWebAclCmd = &cobra.Command{
	Use:   "update-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateWebAclCmd).Standalone()

	waf_updateWebAclCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateWebAclCmd.Flags().String("default-action", "", "A default action for the web ACL, either ALLOW or BLOCK.")
	waf_updateWebAclCmd.Flags().String("updates", "", "An array of updates to make to the [WebACL]().")
	waf_updateWebAclCmd.Flags().String("web-aclid", "", "The `WebACLId` of the [WebACL]() that you want to update.")
	waf_updateWebAclCmd.MarkFlagRequired("change-token")
	waf_updateWebAclCmd.MarkFlagRequired("web-aclid")
	wafCmd.AddCommand(waf_updateWebAclCmd)
}
