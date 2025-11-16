package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listXssMatchSetsCmd = &cobra.Command{
	Use:   "list-xss-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listXssMatchSetsCmd).Standalone()

	wafRegional_listXssMatchSetsCmd.Flags().String("limit", "", "Specifies the number of [XssMatchSet]() objects that you want AWS WAF to return for this request.")
	wafRegional_listXssMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more [XssMatchSet]() objects than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `XssMatchSets`.")
	wafRegionalCmd.AddCommand(wafRegional_listXssMatchSetsCmd)
}
