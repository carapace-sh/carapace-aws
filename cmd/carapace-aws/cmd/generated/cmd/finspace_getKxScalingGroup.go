package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxScalingGroupCmd = &cobra.Command{
	Use:   "get-kx-scaling-group",
	Short: "Retrieves details of a scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxScalingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_getKxScalingGroupCmd).Standalone()

		finspace_getKxScalingGroupCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_getKxScalingGroupCmd.Flags().String("scaling-group-name", "", "A unique identifier for the kdb scaling group.")
		finspace_getKxScalingGroupCmd.MarkFlagRequired("environment-id")
		finspace_getKxScalingGroupCmd.MarkFlagRequired("scaling-group-name")
	})
	finspaceCmd.AddCommand(finspace_getKxScalingGroupCmd)
}
