package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_cancelDetectMitigationActionsTaskCmd = &cobra.Command{
	Use:   "cancel-detect-mitigation-actions-task",
	Short: "Cancels a Device Defender ML Detect mitigation action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_cancelDetectMitigationActionsTaskCmd).Standalone()

	iot_cancelDetectMitigationActionsTaskCmd.Flags().String("task-id", "", "The unique identifier of the task.")
	iot_cancelDetectMitigationActionsTaskCmd.MarkFlagRequired("task-id")
	iotCmd.AddCommand(iot_cancelDetectMitigationActionsTaskCmd)
}
