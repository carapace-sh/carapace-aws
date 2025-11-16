package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetLocationUtilizationCmd = &cobra.Command{
	Use:   "describe-fleet-location-utilization",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves current usage data for a fleet location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetLocationUtilizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeFleetLocationUtilizationCmd).Standalone()

		gamelift_describeFleetLocationUtilizationCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to request location utilization for.")
		gamelift_describeFleetLocationUtilizationCmd.Flags().String("location", "", "The fleet location to retrieve utilization information for.")
		gamelift_describeFleetLocationUtilizationCmd.MarkFlagRequired("fleet-id")
		gamelift_describeFleetLocationUtilizationCmd.MarkFlagRequired("location")
	})
	gameliftCmd.AddCommand(gamelift_describeFleetLocationUtilizationCmd)
}
