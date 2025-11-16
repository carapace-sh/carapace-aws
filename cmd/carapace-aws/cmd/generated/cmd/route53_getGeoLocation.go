package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getGeoLocationCmd = &cobra.Command{
	Use:   "get-geo-location",
	Short: "Gets information about whether a specified geographic location is supported for Amazon Route 53 geolocation resource record sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getGeoLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getGeoLocationCmd).Standalone()

		route53_getGeoLocationCmd.Flags().String("continent-code", "", "For geolocation resource record sets, a two-letter abbreviation that identifies a continent.")
		route53_getGeoLocationCmd.Flags().String("country-code", "", "Amazon Route 53 uses the two-letter country codes that are specified in [ISO standard 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).")
		route53_getGeoLocationCmd.Flags().String("subdivision-code", "", "The code for the subdivision, such as a particular state within the United States.")
	})
	route53Cmd.AddCommand(route53_getGeoLocationCmd)
}
