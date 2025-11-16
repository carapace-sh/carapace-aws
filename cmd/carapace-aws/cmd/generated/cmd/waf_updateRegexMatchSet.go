package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateRegexMatchSetCmd = &cobra.Command{
	Use:   "update-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateRegexMatchSetCmd).Standalone()

	waf_updateRegexMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateRegexMatchSetCmd.Flags().String("regex-match-set-id", "", "The `RegexMatchSetId` of the [RegexMatchSet]() that you want to update.")
	waf_updateRegexMatchSetCmd.Flags().String("updates", "", "An array of `RegexMatchSetUpdate` objects that you want to insert into or delete from a [RegexMatchSet]().")
	waf_updateRegexMatchSetCmd.MarkFlagRequired("change-token")
	waf_updateRegexMatchSetCmd.MarkFlagRequired("regex-match-set-id")
	waf_updateRegexMatchSetCmd.MarkFlagRequired("updates")
	wafCmd.AddCommand(waf_updateRegexMatchSetCmd)
}
