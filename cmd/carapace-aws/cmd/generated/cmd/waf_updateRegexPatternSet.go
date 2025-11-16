package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateRegexPatternSetCmd = &cobra.Command{
	Use:   "update-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateRegexPatternSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_updateRegexPatternSetCmd).Standalone()

		waf_updateRegexPatternSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_updateRegexPatternSetCmd.Flags().String("regex-pattern-set-id", "", "The `RegexPatternSetId` of the [RegexPatternSet]() that you want to update.")
		waf_updateRegexPatternSetCmd.Flags().String("updates", "", "An array of `RegexPatternSetUpdate` objects that you want to insert into or delete from a [RegexPatternSet]().")
		waf_updateRegexPatternSetCmd.MarkFlagRequired("change-token")
		waf_updateRegexPatternSetCmd.MarkFlagRequired("regex-pattern-set-id")
		waf_updateRegexPatternSetCmd.MarkFlagRequired("updates")
	})
	wafCmd.AddCommand(waf_updateRegexPatternSetCmd)
}
