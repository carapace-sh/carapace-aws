package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteXssMatchSetCmd = &cobra.Command{
	Use:   "delete-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteXssMatchSetCmd).Standalone()

		waf_deleteXssMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteXssMatchSetCmd.Flags().String("xss-match-set-id", "", "The `XssMatchSetId` of the [XssMatchSet]() that you want to delete.")
		waf_deleteXssMatchSetCmd.MarkFlagRequired("change-token")
		waf_deleteXssMatchSetCmd.MarkFlagRequired("xss-match-set-id")
	})
	wafCmd.AddCommand(waf_deleteXssMatchSetCmd)
}
