package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listGeoLocationsCmd = &cobra.Command{
	Use:   "list-geo-locations",
	Short: "Retrieves a list of supported geographic locations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listGeoLocationsCmd).Standalone()

	route53_listGeoLocationsCmd.Flags().String("max-items", "", "(Optional) The maximum number of geolocations to be included in the response body for this request.")
	route53_listGeoLocationsCmd.Flags().String("start-continent-code", "", "The code for the continent with which you want to start listing locations that Amazon Route 53 supports for geolocation.")
	route53_listGeoLocationsCmd.Flags().String("start-country-code", "", "The code for the country with which you want to start listing locations that Amazon Route 53 supports for geolocation.")
	route53_listGeoLocationsCmd.Flags().String("start-subdivision-code", "", "The code for the state of the United States with which you want to start listing locations that Amazon Route 53 supports for geolocation.")
	route53Cmd.AddCommand(route53_listGeoLocationsCmd)
}
