package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getGeoMatchSetCmd = &cobra.Command{
	Use:   "get-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getGeoMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getGeoMatchSetCmd).Standalone()

		wafRegional_getGeoMatchSetCmd.Flags().String("geo-match-set-id", "", "The `GeoMatchSetId` of the [GeoMatchSet]() that you want to get.")
		wafRegional_getGeoMatchSetCmd.MarkFlagRequired("geo-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_getGeoMatchSetCmd)
}
