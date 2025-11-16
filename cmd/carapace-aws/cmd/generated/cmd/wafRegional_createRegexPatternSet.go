package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createRegexPatternSetCmd = &cobra.Command{
	Use:   "create-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createRegexPatternSetCmd).Standalone()

	wafRegional_createRegexPatternSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_createRegexPatternSetCmd.Flags().String("name", "", "A friendly name or description of the [RegexPatternSet]().")
	wafRegional_createRegexPatternSetCmd.MarkFlagRequired("change-token")
	wafRegional_createRegexPatternSetCmd.MarkFlagRequired("name")
	wafRegionalCmd.AddCommand(wafRegional_createRegexPatternSetCmd)
}
