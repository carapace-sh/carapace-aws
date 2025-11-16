package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_batchDeleteScheduledActionCmd = &cobra.Command{
	Use:   "batch-delete-scheduled-action",
	Short: "Deletes one or more scheduled actions for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_batchDeleteScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_batchDeleteScheduledActionCmd).Standalone()

		autoscaling_batchDeleteScheduledActionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_batchDeleteScheduledActionCmd.Flags().String("scheduled-action-names", "", "The names of the scheduled actions to delete.")
		autoscaling_batchDeleteScheduledActionCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_batchDeleteScheduledActionCmd.MarkFlagRequired("scheduled-action-names")
	})
	autoscalingCmd.AddCommand(autoscaling_batchDeleteScheduledActionCmd)
}
