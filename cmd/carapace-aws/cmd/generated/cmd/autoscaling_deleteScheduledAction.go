package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_deleteScheduledActionCmd = &cobra.Command{
	Use:   "delete-scheduled-action",
	Short: "Deletes the specified scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_deleteScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_deleteScheduledActionCmd).Standalone()

		autoscaling_deleteScheduledActionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_deleteScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the action to delete.")
		autoscaling_deleteScheduledActionCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_deleteScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	})
	autoscalingCmd.AddCommand(autoscaling_deleteScheduledActionCmd)
}
