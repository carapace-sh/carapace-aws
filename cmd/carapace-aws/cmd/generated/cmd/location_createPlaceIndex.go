package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_createPlaceIndexCmd = &cobra.Command{
	Use:   "create-place-index",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_createPlaceIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_createPlaceIndexCmd).Standalone()

		location_createPlaceIndexCmd.Flags().String("data-source", "", "Specifies the geospatial data provider for the new place index.")
		location_createPlaceIndexCmd.Flags().String("data-source-configuration", "", "Specifies the data storage option requesting Places.")
		location_createPlaceIndexCmd.Flags().String("description", "", "The optional description for the place index resource.")
		location_createPlaceIndexCmd.Flags().String("index-name", "", "The name of the place index resource.")
		location_createPlaceIndexCmd.Flags().String("pricing-plan", "", "No longer used.")
		location_createPlaceIndexCmd.Flags().String("tags", "", "Applies one or more tags to the place index resource.")
		location_createPlaceIndexCmd.MarkFlagRequired("data-source")
		location_createPlaceIndexCmd.MarkFlagRequired("index-name")
	})
	locationCmd.AddCommand(location_createPlaceIndexCmd)
}
