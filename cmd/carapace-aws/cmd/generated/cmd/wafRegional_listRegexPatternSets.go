package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listRegexPatternSetsCmd = &cobra.Command{
	Use:   "list-regex-pattern-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listRegexPatternSetsCmd).Standalone()

	wafRegional_listRegexPatternSetsCmd.Flags().String("limit", "", "Specifies the number of `RegexPatternSet` objects that you want AWS WAF to return for this request.")
	wafRegional_listRegexPatternSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `RegexPatternSet` objects than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `RegexPatternSet` objects.")
	wafRegionalCmd.AddCommand(wafRegional_listRegexPatternSetsCmd)
}
