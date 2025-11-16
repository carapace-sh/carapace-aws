package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteRegexMatchSetCmd = &cobra.Command{
	Use:   "delete-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteRegexMatchSetCmd).Standalone()

	wafRegional_deleteRegexMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_deleteRegexMatchSetCmd.Flags().String("regex-match-set-id", "", "The `RegexMatchSetId` of the [RegexMatchSet]() that you want to delete.")
	wafRegional_deleteRegexMatchSetCmd.MarkFlagRequired("change-token")
	wafRegional_deleteRegexMatchSetCmd.MarkFlagRequired("regex-match-set-id")
	wafRegionalCmd.AddCommand(wafRegional_deleteRegexMatchSetCmd)
}
