package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createRegexMatchSetCmd = &cobra.Command{
	Use:   "create-regex-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createRegexMatchSetCmd).Standalone()

	waf_createRegexMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_createRegexMatchSetCmd.Flags().String("name", "", "A friendly name or description of the [RegexMatchSet]().")
	waf_createRegexMatchSetCmd.MarkFlagRequired("change-token")
	waf_createRegexMatchSetCmd.MarkFlagRequired("name")
	wafCmd.AddCommand(waf_createRegexMatchSetCmd)
}
