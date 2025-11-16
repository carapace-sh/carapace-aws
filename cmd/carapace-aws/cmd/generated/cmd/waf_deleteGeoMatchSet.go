package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteGeoMatchSetCmd = &cobra.Command{
	Use:   "delete-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteGeoMatchSetCmd).Standalone()

	waf_deleteGeoMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_deleteGeoMatchSetCmd.Flags().String("geo-match-set-id", "", "The `GeoMatchSetID` of the [GeoMatchSet]() that you want to delete.")
	waf_deleteGeoMatchSetCmd.MarkFlagRequired("change-token")
	waf_deleteGeoMatchSetCmd.MarkFlagRequired("geo-match-set-id")
	wafCmd.AddCommand(waf_deleteGeoMatchSetCmd)
}
