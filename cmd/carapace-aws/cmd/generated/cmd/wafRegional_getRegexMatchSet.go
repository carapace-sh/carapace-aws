package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getRegexMatchSetCmd = &cobra.Command{
	Use:   "get-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getRegexMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getRegexMatchSetCmd).Standalone()

		wafRegional_getRegexMatchSetCmd.Flags().String("regex-match-set-id", "", "The `RegexMatchSetId` of the [RegexMatchSet]() that you want to get.")
		wafRegional_getRegexMatchSetCmd.MarkFlagRequired("regex-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_getRegexMatchSetCmd)
}
