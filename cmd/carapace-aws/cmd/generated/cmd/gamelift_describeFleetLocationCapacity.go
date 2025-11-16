package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetLocationCapacityCmd = &cobra.Command{
	Use:   "describe-fleet-location-capacity",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRetrieves the resource capacity settings for a fleet location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetLocationCapacityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeFleetLocationCapacityCmd).Standalone()

		gamelift_describeFleetLocationCapacityCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to request location capacity for.")
		gamelift_describeFleetLocationCapacityCmd.Flags().String("location", "", "The fleet location to retrieve capacity information for.")
		gamelift_describeFleetLocationCapacityCmd.MarkFlagRequired("fleet-id")
		gamelift_describeFleetLocationCapacityCmd.MarkFlagRequired("location")
	})
	gameliftCmd.AddCommand(gamelift_describeFleetLocationCapacityCmd)
}
