package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getMapStyleDescriptorCmd = &cobra.Command{
	Use:   "get-map-style-descriptor",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getMapStyleDescriptorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_getMapStyleDescriptorCmd).Standalone()

		location_getMapStyleDescriptorCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
		location_getMapStyleDescriptorCmd.Flags().String("map-name", "", "The map resource to retrieve the style descriptor from.")
		location_getMapStyleDescriptorCmd.MarkFlagRequired("map-name")
	})
	locationCmd.AddCommand(location_getMapStyleDescriptorCmd)
}
