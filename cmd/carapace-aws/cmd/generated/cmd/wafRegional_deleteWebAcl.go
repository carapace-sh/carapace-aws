package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteWebAclCmd = &cobra.Command{
	Use:   "delete-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteWebAclCmd).Standalone()

	wafRegional_deleteWebAclCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_deleteWebAclCmd.Flags().String("web-aclid", "", "The `WebACLId` of the [WebACL]() that you want to delete.")
	wafRegional_deleteWebAclCmd.MarkFlagRequired("change-token")
	wafRegional_deleteWebAclCmd.MarkFlagRequired("web-aclid")
	wafRegionalCmd.AddCommand(wafRegional_deleteWebAclCmd)
}
