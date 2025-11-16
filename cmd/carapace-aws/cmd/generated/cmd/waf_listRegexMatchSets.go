package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listRegexMatchSetsCmd = &cobra.Command{
	Use:   "list-regex-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listRegexMatchSetsCmd).Standalone()

	waf_listRegexMatchSetsCmd.Flags().String("limit", "", "Specifies the number of `RegexMatchSet` objects that you want AWS WAF to return for this request.")
	waf_listRegexMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `RegexMatchSet` objects than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `ByteMatchSets`.")
	wafCmd.AddCommand(waf_listRegexMatchSetsCmd)
}
