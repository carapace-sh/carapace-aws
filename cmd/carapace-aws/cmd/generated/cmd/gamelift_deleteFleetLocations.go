package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteFleetLocationsCmd = &cobra.Command{
	Use:   "delete-fleet-locations",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRemoves locations from a multi-location fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteFleetLocationsCmd).Standalone()

	gamelift_deleteFleetLocationsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to delete locations for.")
	gamelift_deleteFleetLocationsCmd.Flags().String("locations", "", "The list of fleet locations to delete.")
	gamelift_deleteFleetLocationsCmd.MarkFlagRequired("fleet-id")
	gamelift_deleteFleetLocationsCmd.MarkFlagRequired("locations")
	gameliftCmd.AddCommand(gamelift_deleteFleetLocationsCmd)
}
