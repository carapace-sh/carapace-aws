package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_updatePlaceIndexCmd = &cobra.Command{
	Use:   "update-place-index",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_updatePlaceIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_updatePlaceIndexCmd).Standalone()

		location_updatePlaceIndexCmd.Flags().String("data-source-configuration", "", "Updates the data storage option for the place index resource.")
		location_updatePlaceIndexCmd.Flags().String("description", "", "Updates the description for the place index resource.")
		location_updatePlaceIndexCmd.Flags().String("index-name", "", "The name of the place index resource to update.")
		location_updatePlaceIndexCmd.Flags().String("pricing-plan", "", "No longer used.")
		location_updatePlaceIndexCmd.MarkFlagRequired("index-name")
	})
	locationCmd.AddCommand(location_updatePlaceIndexCmd)
}
