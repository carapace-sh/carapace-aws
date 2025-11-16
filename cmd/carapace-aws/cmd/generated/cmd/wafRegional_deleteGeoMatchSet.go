package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteGeoMatchSetCmd = &cobra.Command{
	Use:   "delete-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteGeoMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_deleteGeoMatchSetCmd).Standalone()

		wafRegional_deleteGeoMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_deleteGeoMatchSetCmd.Flags().String("geo-match-set-id", "", "The `GeoMatchSetID` of the [GeoMatchSet]() that you want to delete.")
		wafRegional_deleteGeoMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_deleteGeoMatchSetCmd.MarkFlagRequired("geo-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_deleteGeoMatchSetCmd)
}
