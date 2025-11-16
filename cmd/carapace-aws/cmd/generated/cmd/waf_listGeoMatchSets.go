package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listGeoMatchSetsCmd = &cobra.Command{
	Use:   "list-geo-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listGeoMatchSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_listGeoMatchSetsCmd).Standalone()

		waf_listGeoMatchSetsCmd.Flags().String("limit", "", "Specifies the number of `GeoMatchSet` objects that you want AWS WAF to return for this request.")
		waf_listGeoMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `GeoMatchSet`s than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `GeoMatchSet` objects.")
	})
	wafCmd.AddCommand(waf_listGeoMatchSetsCmd)
}
