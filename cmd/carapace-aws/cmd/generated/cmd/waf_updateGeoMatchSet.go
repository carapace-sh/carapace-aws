package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateGeoMatchSetCmd = &cobra.Command{
	Use:   "update-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateGeoMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_updateGeoMatchSetCmd).Standalone()

		waf_updateGeoMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_updateGeoMatchSetCmd.Flags().String("geo-match-set-id", "", "The `GeoMatchSetId` of the [GeoMatchSet]() that you want to update.")
		waf_updateGeoMatchSetCmd.Flags().String("updates", "", "An array of `GeoMatchSetUpdate` objects that you want to insert into or delete from an [GeoMatchSet]().")
		waf_updateGeoMatchSetCmd.MarkFlagRequired("change-token")
		waf_updateGeoMatchSetCmd.MarkFlagRequired("geo-match-set-id")
		waf_updateGeoMatchSetCmd.MarkFlagRequired("updates")
	})
	wafCmd.AddCommand(waf_updateGeoMatchSetCmd)
}
