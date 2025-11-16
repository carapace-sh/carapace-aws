package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createFleetLocationsCmd = &cobra.Command{
	Use:   "create-fleet-locations",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nAdds remote locations to an EC2 and begins populating the new locations with instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createFleetLocationsCmd).Standalone()

	gamelift_createFleetLocationsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to add locations to.")
	gamelift_createFleetLocationsCmd.Flags().String("locations", "", "A list of locations to deploy additional instances to and manage as part of the fleet.")
	gamelift_createFleetLocationsCmd.MarkFlagRequired("fleet-id")
	gamelift_createFleetLocationsCmd.MarkFlagRequired("locations")
	gameliftCmd.AddCommand(gamelift_createFleetLocationsCmd)
}
