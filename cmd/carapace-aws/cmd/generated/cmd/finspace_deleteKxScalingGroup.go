package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxScalingGroupCmd = &cobra.Command{
	Use:   "delete-kx-scaling-group",
	Short: "Deletes the specified scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxScalingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_deleteKxScalingGroupCmd).Standalone()

		finspace_deleteKxScalingGroupCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspace_deleteKxScalingGroupCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, from where you want to delete the dataview.")
		finspace_deleteKxScalingGroupCmd.Flags().String("scaling-group-name", "", "A unique identifier for the kdb scaling group.")
		finspace_deleteKxScalingGroupCmd.MarkFlagRequired("environment-id")
		finspace_deleteKxScalingGroupCmd.MarkFlagRequired("scaling-group-name")
	})
	finspaceCmd.AddCommand(finspace_deleteKxScalingGroupCmd)
}
