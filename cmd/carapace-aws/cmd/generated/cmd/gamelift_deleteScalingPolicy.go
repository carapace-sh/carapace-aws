package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteScalingPolicyCmd = &cobra.Command{
	Use:   "delete-scaling-policy",
	Short: "**This API works with the following fleet types:** EC2\n\nDeletes a fleet scaling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteScalingPolicyCmd).Standalone()

	gamelift_deleteScalingPolicyCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to be deleted.")
	gamelift_deleteScalingPolicyCmd.Flags().String("name", "", "A descriptive label that is associated with a fleet's scaling policy.")
	gamelift_deleteScalingPolicyCmd.MarkFlagRequired("fleet-id")
	gamelift_deleteScalingPolicyCmd.MarkFlagRequired("name")
	gameliftCmd.AddCommand(gamelift_deleteScalingPolicyCmd)
}
