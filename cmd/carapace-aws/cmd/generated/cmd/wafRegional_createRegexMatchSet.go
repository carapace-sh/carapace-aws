package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createRegexMatchSetCmd = &cobra.Command{
	Use:   "create-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createRegexMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createRegexMatchSetCmd).Standalone()

		wafRegional_createRegexMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createRegexMatchSetCmd.Flags().String("name", "", "A friendly name or description of the [RegexMatchSet]().")
		wafRegional_createRegexMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_createRegexMatchSetCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createRegexMatchSetCmd)
}
