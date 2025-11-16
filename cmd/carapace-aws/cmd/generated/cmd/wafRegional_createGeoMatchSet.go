package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createGeoMatchSetCmd = &cobra.Command{
	Use:   "create-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createGeoMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createGeoMatchSetCmd).Standalone()

		wafRegional_createGeoMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createGeoMatchSetCmd.Flags().String("name", "", "A friendly name or description of the [GeoMatchSet]().")
		wafRegional_createGeoMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_createGeoMatchSetCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createGeoMatchSetCmd)
}
