package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteRegexMatchSetCmd = &cobra.Command{
	Use:   "delete-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteRegexMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteRegexMatchSetCmd).Standalone()

		waf_deleteRegexMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteRegexMatchSetCmd.Flags().String("regex-match-set-id", "", "The `RegexMatchSetId` of the [RegexMatchSet]() that you want to delete.")
		waf_deleteRegexMatchSetCmd.MarkFlagRequired("change-token")
		waf_deleteRegexMatchSetCmd.MarkFlagRequired("regex-match-set-id")
	})
	wafCmd.AddCommand(waf_deleteRegexMatchSetCmd)
}
