package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteXssMatchSetCmd = &cobra.Command{
	Use:   "delete-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteXssMatchSetCmd).Standalone()

		wafRegional_deleteXssMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteXssMatchSetCmd.Flags().String("xss-match-set-id", "", "The `XssMatchSetId` of the [XssMatchSet]() that you want to delete.")
		wafRegional_deleteXssMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_deleteXssMatchSetCmd.MarkFlagRequired("xss-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteXssMatchSetCmd)
}
