package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createXssMatchSetCmd = &cobra.Command{
	Use:   "create-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createXssMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createXssMatchSetCmd).Standalone()

		wafRegional_createXssMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createXssMatchSetCmd.Flags().String("name", "", "A friendly name or description for the [XssMatchSet]() that you're creating.")
		wafRegional_createXssMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_createXssMatchSetCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createXssMatchSetCmd)
}
