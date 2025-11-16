package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createGeoMatchSetCmd = &cobra.Command{
	Use:   "create-geo-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createGeoMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createGeoMatchSetCmd).Standalone()

		waf_createGeoMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createGeoMatchSetCmd.Flags().String("name", "", "A friendly name or description of the [GeoMatchSet]().")
		waf_createGeoMatchSetCmd.MarkFlagRequired("change-token")
		waf_createGeoMatchSetCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createGeoMatchSetCmd)
}
