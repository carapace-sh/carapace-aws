package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateRegexPatternSetCmd = &cobra.Command{
	Use:   "update-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateRegexPatternSetCmd).Standalone()

	wafRegional_updateRegexPatternSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateRegexPatternSetCmd.Flags().String("regex-pattern-set-id", "", "The `RegexPatternSetId` of the [RegexPatternSet]() that you want to update.")
	wafRegional_updateRegexPatternSetCmd.Flags().String("updates", "", "An array of `RegexPatternSetUpdate` objects that you want to insert into or delete from a [RegexPatternSet]().")
	wafRegional_updateRegexPatternSetCmd.MarkFlagRequired("change-token")
	wafRegional_updateRegexPatternSetCmd.MarkFlagRequired("regex-pattern-set-id")
	wafRegional_updateRegexPatternSetCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateRegexPatternSetCmd)
}
