package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteRegexPatternSetCmd = &cobra.Command{
	Use:   "delete-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteRegexPatternSetCmd).Standalone()

	waf_deleteRegexPatternSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_deleteRegexPatternSetCmd.Flags().String("regex-pattern-set-id", "", "The `RegexPatternSetId` of the [RegexPatternSet]() that you want to delete.")
	waf_deleteRegexPatternSetCmd.MarkFlagRequired("change-token")
	waf_deleteRegexPatternSetCmd.MarkFlagRequired("regex-pattern-set-id")
	wafCmd.AddCommand(waf_deleteRegexPatternSetCmd)
}
