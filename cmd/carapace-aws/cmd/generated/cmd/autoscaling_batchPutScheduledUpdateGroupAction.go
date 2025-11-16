package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_batchPutScheduledUpdateGroupActionCmd = &cobra.Command{
	Use:   "batch-put-scheduled-update-group-action",
	Short: "Creates or updates one or more scheduled scaling actions for an Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_batchPutScheduledUpdateGroupActionCmd).Standalone()

	autoscaling_batchPutScheduledUpdateGroupActionCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_batchPutScheduledUpdateGroupActionCmd.Flags().String("scheduled-update-group-actions", "", "One or more scheduled actions.")
	autoscaling_batchPutScheduledUpdateGroupActionCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_batchPutScheduledUpdateGroupActionCmd.MarkFlagRequired("scheduled-update-group-actions")
	autoscalingCmd.AddCommand(autoscaling_batchPutScheduledUpdateGroupActionCmd)
}
