package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetUtilizationCmd = &cobra.Command{
	Use:   "describe-fleet-utilization",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRetrieves utilization statistics for one or more fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetUtilizationCmd).Standalone()

	gamelift_describeFleetUtilizationCmd.Flags().String("fleet-ids", "", "A unique identifier for the fleet to retrieve utilization data for.")
	gamelift_describeFleetUtilizationCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_describeFleetUtilizationCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gameliftCmd.AddCommand(gamelift_describeFleetUtilizationCmd)
}
