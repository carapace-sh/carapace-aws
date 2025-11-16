package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listDevicePositionsCmd = &cobra.Command{
	Use:   "list-device-positions",
	Short: "A batch request to retrieve all device positions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listDevicePositionsCmd).Standalone()

	location_listDevicePositionsCmd.Flags().String("filter-geometry", "", "The geometry used to filter device positions.")
	location_listDevicePositionsCmd.Flags().String("max-results", "", "An optional limit for the number of entries returned in a single call.")
	location_listDevicePositionsCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	location_listDevicePositionsCmd.Flags().String("tracker-name", "", "The tracker resource containing the requested devices.")
	location_listDevicePositionsCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_listDevicePositionsCmd)
}
