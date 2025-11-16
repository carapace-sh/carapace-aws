package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getXssMatchSetCmd = &cobra.Command{
	Use:   "get-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_getXssMatchSetCmd).Standalone()

		waf_getXssMatchSetCmd.Flags().String("xss-match-set-id", "", "The `XssMatchSetId` of the [XssMatchSet]() that you want to get.")
		waf_getXssMatchSetCmd.MarkFlagRequired("xss-match-set-id")
	})
	wafCmd.AddCommand(waf_getXssMatchSetCmd)
}
