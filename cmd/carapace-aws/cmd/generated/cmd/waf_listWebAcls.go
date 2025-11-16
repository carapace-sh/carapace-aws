package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listWebAclsCmd = &cobra.Command{
	Use:   "list-web-acls",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listWebAclsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_listWebAclsCmd).Standalone()

		waf_listWebAclsCmd.Flags().String("limit", "", "Specifies the number of `WebACL` objects that you want AWS WAF to return for this request.")
		waf_listWebAclsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `WebACL` objects than the number that you specify for `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `WebACL` objects.")
	})
	wafCmd.AddCommand(waf_listWebAclsCmd)
}
