package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeDetectMitigationActionsTaskCmd = &cobra.Command{
	Use:   "describe-detect-mitigation-actions-task",
	Short: "Gets information about a Device Defender ML Detect mitigation action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeDetectMitigationActionsTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeDetectMitigationActionsTaskCmd).Standalone()

		iot_describeDetectMitigationActionsTaskCmd.Flags().String("task-id", "", "The unique identifier of the task.")
		iot_describeDetectMitigationActionsTaskCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_describeDetectMitigationActionsTaskCmd)
}
