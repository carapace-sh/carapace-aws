package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteRegexPatternSetCmd = &cobra.Command{
	Use:   "delete-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteRegexPatternSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteRegexPatternSetCmd).Standalone()

		wafRegional_deleteRegexPatternSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteRegexPatternSetCmd.Flags().String("regex-pattern-set-id", "", "The `RegexPatternSetId` of the [RegexPatternSet]() that you want to delete.")
		wafRegional_deleteRegexPatternSetCmd.MarkFlagRequired("change-token")
		wafRegional_deleteRegexPatternSetCmd.MarkFlagRequired("regex-pattern-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteRegexPatternSetCmd)
}
