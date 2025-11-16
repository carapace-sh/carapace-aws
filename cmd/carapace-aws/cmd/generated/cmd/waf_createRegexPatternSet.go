package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createRegexPatternSetCmd = &cobra.Command{
	Use:   "create-regex-pattern-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createRegexPatternSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createRegexPatternSetCmd).Standalone()

		waf_createRegexPatternSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createRegexPatternSetCmd.Flags().String("name", "", "A friendly name or description of the [RegexPatternSet]().")
		waf_createRegexPatternSetCmd.MarkFlagRequired("change-token")
		waf_createRegexPatternSetCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createRegexPatternSetCmd)
}
