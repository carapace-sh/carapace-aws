package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getXssMatchSetCmd = &cobra.Command{
	Use:   "get-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getXssMatchSetCmd).Standalone()

		wafRegional_getXssMatchSetCmd.Flags().String("xss-match-set-id", "", "The `XssMatchSetId` of the [XssMatchSet]() that you want to get.")
		wafRegional_getXssMatchSetCmd.MarkFlagRequired("xss-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_getXssMatchSetCmd)
}
