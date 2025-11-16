package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listGeoMatchSetsCmd = &cobra.Command{
	Use:   "list-geo-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listGeoMatchSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_listGeoMatchSetsCmd).Standalone()

		wafRegional_listGeoMatchSetsCmd.Flags().String("limit", "", "Specifies the number of `GeoMatchSet` objects that you want AWS WAF to return for this request.")
		wafRegional_listGeoMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `GeoMatchSet`s than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `GeoMatchSet` objects.")
	})
	wafRegionalCmd.AddCommand(wafRegional_listGeoMatchSetsCmd)
}
