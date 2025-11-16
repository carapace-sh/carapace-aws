package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateXssMatchSetCmd = &cobra.Command{
	Use:   "update-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_updateXssMatchSetCmd).Standalone()

		wafRegional_updateXssMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_updateXssMatchSetCmd.Flags().String("updates", "", "An array of `XssMatchSetUpdate` objects that you want to insert into or delete from an [XssMatchSet]().")
		wafRegional_updateXssMatchSetCmd.Flags().String("xss-match-set-id", "", "The `XssMatchSetId` of the `XssMatchSet` that you want to update.")
		wafRegional_updateXssMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_updateXssMatchSetCmd.MarkFlagRequired("updates")
		wafRegional_updateXssMatchSetCmd.MarkFlagRequired("xss-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_updateXssMatchSetCmd)
}
