package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateWebAclCmd = &cobra.Command{
	Use:   "update-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_updateWebAclCmd).Standalone()

		wafRegional_updateWebAclCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_updateWebAclCmd.Flags().String("default-action", "", "A default action for the web ACL, either ALLOW or BLOCK.")
		wafRegional_updateWebAclCmd.Flags().String("updates", "", "An array of updates to make to the [WebACL]().")
		wafRegional_updateWebAclCmd.Flags().String("web-aclid", "", "The `WebACLId` of the [WebACL]() that you want to update.")
		wafRegional_updateWebAclCmd.MarkFlagRequired("change-token")
		wafRegional_updateWebAclCmd.MarkFlagRequired("web-aclid")
	})
	wafRegionalCmd.AddCommand(wafRegional_updateWebAclCmd)
}
