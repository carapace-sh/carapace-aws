package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateFleetCapacityCmd = &cobra.Command{
	Use:   "update-fleet-capacity",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nUpdates capacity settings for a managed EC2 fleet or managed container fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateFleetCapacityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_updateFleetCapacityCmd).Standalone()

		gamelift_updateFleetCapacityCmd.Flags().String("desired-instances", "", "The number of Amazon EC2 instances you want to maintain in the specified fleet location.")
		gamelift_updateFleetCapacityCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to update capacity settings for.")
		gamelift_updateFleetCapacityCmd.Flags().String("location", "", "The name of a remote location to update fleet capacity settings for, in the form of an Amazon Web Services Region code such as `us-west-2`.")
		gamelift_updateFleetCapacityCmd.Flags().String("max-size", "", "The maximum number of instances that are allowed in the specified fleet location.")
		gamelift_updateFleetCapacityCmd.Flags().String("min-size", "", "The minimum number of instances that are allowed in the specified fleet location.")
		gamelift_updateFleetCapacityCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_updateFleetCapacityCmd)
}
