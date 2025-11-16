package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getGeoMatchSetCmd = &cobra.Command{
	Use:   "get-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getGeoMatchSetCmd).Standalone()

	waf_getGeoMatchSetCmd.Flags().String("geo-match-set-id", "", "The `GeoMatchSetId` of the [GeoMatchSet]() that you want to get.")
	waf_getGeoMatchSetCmd.MarkFlagRequired("geo-match-set-id")
	wafCmd.AddCommand(waf_getGeoMatchSetCmd)
}
