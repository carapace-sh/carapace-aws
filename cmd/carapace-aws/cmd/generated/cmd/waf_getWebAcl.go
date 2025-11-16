package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getWebAclCmd = &cobra.Command{
	Use:   "get-web-acl",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getWebAclCmd).Standalone()

	waf_getWebAclCmd.Flags().String("web-aclid", "", "The `WebACLId` of the [WebACL]() that you want to get.")
	waf_getWebAclCmd.MarkFlagRequired("web-aclid")
	wafCmd.AddCommand(waf_getWebAclCmd)
}
