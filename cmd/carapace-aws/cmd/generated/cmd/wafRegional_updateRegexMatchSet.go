package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateRegexMatchSetCmd = &cobra.Command{
	Use:   "update-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateRegexMatchSetCmd).Standalone()

	wafRegional_updateRegexMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateRegexMatchSetCmd.Flags().String("regex-match-set-id", "", "The `RegexMatchSetId` of the [RegexMatchSet]() that you want to update.")
	wafRegional_updateRegexMatchSetCmd.Flags().String("updates", "", "An array of `RegexMatchSetUpdate` objects that you want to insert into or delete from a [RegexMatchSet]().")
	wafRegional_updateRegexMatchSetCmd.MarkFlagRequired("change-token")
	wafRegional_updateRegexMatchSetCmd.MarkFlagRequired("regex-match-set-id")
	wafRegional_updateRegexMatchSetCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateRegexMatchSetCmd)
}
