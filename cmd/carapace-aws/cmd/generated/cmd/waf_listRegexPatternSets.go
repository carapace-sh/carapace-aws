package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listRegexPatternSetsCmd = &cobra.Command{
	Use:   "list-regex-pattern-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listRegexPatternSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_listRegexPatternSetsCmd).Standalone()

		waf_listRegexPatternSetsCmd.Flags().String("limit", "", "Specifies the number of `RegexPatternSet` objects that you want AWS WAF to return for this request.")
		waf_listRegexPatternSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `RegexPatternSet` objects than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `RegexPatternSet` objects.")
	})
	wafCmd.AddCommand(waf_listRegexPatternSetsCmd)
}
