package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listRateBasedRulesCmd = &cobra.Command{
	Use:   "list-rate-based-rules",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listRateBasedRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_listRateBasedRulesCmd).Standalone()

		wafRegional_listRateBasedRulesCmd.Flags().String("limit", "", "Specifies the number of `Rules` that you want AWS WAF to return for this request.")
		wafRegional_listRateBasedRulesCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `Rules` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `Rules`.")
	})
	wafRegionalCmd.AddCommand(wafRegional_listRateBasedRulesCmd)
}
