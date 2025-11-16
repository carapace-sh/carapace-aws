package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getRegexPatternSetCmd = &cobra.Command{
	Use:   "get-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getRegexPatternSetCmd).Standalone()

	wafRegional_getRegexPatternSetCmd.Flags().String("regex-pattern-set-id", "", "The `RegexPatternSetId` of the [RegexPatternSet]() that you want to get.")
	wafRegional_getRegexPatternSetCmd.MarkFlagRequired("regex-pattern-set-id")
	wafRegionalCmd.AddCommand(wafRegional_getRegexPatternSetCmd)
}
