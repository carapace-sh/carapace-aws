package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteWarmPoolCmd = &cobra.Command{
	Use:   "delete-warm-pool",
	Short: "Deletes the warm pool for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteWarmPoolCmd).Standalone()

	autoscaling_deleteWarmPoolCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_deleteWarmPoolCmd.Flags().String("force-delete", "", "Specifies that the warm pool is to be deleted along with all of its associated instances, without waiting for all instances to be terminated.")
	autoscaling_deleteWarmPoolCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_deleteWarmPoolCmd)
}
