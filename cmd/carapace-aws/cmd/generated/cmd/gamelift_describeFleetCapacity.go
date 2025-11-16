package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetCapacityCmd = &cobra.Command{
	Use:   "describe-fleet-capacity",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRetrieves the resource capacity settings for one or more fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetCapacityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeFleetCapacityCmd).Standalone()

		gamelift_describeFleetCapacityCmd.Flags().String("fleet-ids", "", "A unique identifier for the fleet to retrieve capacity information for.")
		gamelift_describeFleetCapacityCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_describeFleetCapacityCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_describeFleetCapacityCmd)
}
