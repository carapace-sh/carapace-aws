package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_createMapCmd = &cobra.Command{
	Use:   "create-map",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_createMapCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_createMapCmd).Standalone()

		location_createMapCmd.Flags().String("configuration", "", "Specifies the `MapConfiguration`, including the map style, for the map resource that you create.")
		location_createMapCmd.Flags().String("description", "", "An optional description for the map resource.")
		location_createMapCmd.Flags().String("map-name", "", "The name for the map resource.")
		location_createMapCmd.Flags().String("pricing-plan", "", "No longer used.")
		location_createMapCmd.Flags().String("tags", "", "Applies one or more tags to the map resource.")
		location_createMapCmd.MarkFlagRequired("configuration")
		location_createMapCmd.MarkFlagRequired("map-name")
	})
	locationCmd.AddCommand(location_createMapCmd)
}
