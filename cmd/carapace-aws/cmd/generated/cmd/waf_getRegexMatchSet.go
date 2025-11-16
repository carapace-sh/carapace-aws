package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getRegexMatchSetCmd = &cobra.Command{
	Use:   "get-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getRegexMatchSetCmd).Standalone()

	waf_getRegexMatchSetCmd.Flags().String("regex-match-set-id", "", "The `RegexMatchSetId` of the [RegexMatchSet]() that you want to get.")
	waf_getRegexMatchSetCmd.MarkFlagRequired("regex-match-set-id")
	wafCmd.AddCommand(waf_getRegexMatchSetCmd)
}
