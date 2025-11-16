package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listWebAclsCmd = &cobra.Command{
	Use:   "list-web-acls",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listWebAclsCmd).Standalone()

	wafRegional_listWebAclsCmd.Flags().String("limit", "", "Specifies the number of `WebACL` objects that you want AWS WAF to return for this request.")
	wafRegional_listWebAclsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `WebACL` objects than the number that you specify for `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `WebACL` objects.")
	wafRegionalCmd.AddCommand(wafRegional_listWebAclsCmd)
}
