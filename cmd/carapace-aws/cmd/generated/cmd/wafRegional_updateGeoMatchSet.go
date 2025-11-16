package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateGeoMatchSetCmd = &cobra.Command{
	Use:   "update-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateGeoMatchSetCmd).Standalone()

	wafRegional_updateGeoMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateGeoMatchSetCmd.Flags().String("geo-match-set-id", "", "The `GeoMatchSetId` of the [GeoMatchSet]() that you want to update.")
	wafRegional_updateGeoMatchSetCmd.Flags().String("updates", "", "An array of `GeoMatchSetUpdate` objects that you want to insert into or delete from an [GeoMatchSet]().")
	wafRegional_updateGeoMatchSetCmd.MarkFlagRequired("change-token")
	wafRegional_updateGeoMatchSetCmd.MarkFlagRequired("geo-match-set-id")
	wafRegional_updateGeoMatchSetCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateGeoMatchSetCmd)
}
