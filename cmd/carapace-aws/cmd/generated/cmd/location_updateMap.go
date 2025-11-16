package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_updateMapCmd = &cobra.Command{
	Use:   "update-map",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_updateMapCmd).Standalone()

	location_updateMapCmd.Flags().String("configuration-update", "", "Updates the parts of the map configuration that can be updated, including the political view.")
	location_updateMapCmd.Flags().String("description", "", "Updates the description for the map resource.")
	location_updateMapCmd.Flags().String("map-name", "", "The name of the map resource to update.")
	location_updateMapCmd.Flags().String("pricing-plan", "", "No longer used.")
	location_updateMapCmd.MarkFlagRequired("map-name")
	locationCmd.AddCommand(location_updateMapCmd)
}
