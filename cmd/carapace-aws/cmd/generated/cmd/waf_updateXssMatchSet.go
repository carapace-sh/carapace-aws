package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateXssMatchSetCmd = &cobra.Command{
	Use:   "update-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_updateXssMatchSetCmd).Standalone()

		waf_updateXssMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_updateXssMatchSetCmd.Flags().String("updates", "", "An array of `XssMatchSetUpdate` objects that you want to insert into or delete from an [XssMatchSet]().")
		waf_updateXssMatchSetCmd.Flags().String("xss-match-set-id", "", "The `XssMatchSetId` of the `XssMatchSet` that you want to update.")
		waf_updateXssMatchSetCmd.MarkFlagRequired("change-token")
		waf_updateXssMatchSetCmd.MarkFlagRequired("updates")
		waf_updateXssMatchSetCmd.MarkFlagRequired("xss-match-set-id")
	})
	wafCmd.AddCommand(waf_updateXssMatchSetCmd)
}
