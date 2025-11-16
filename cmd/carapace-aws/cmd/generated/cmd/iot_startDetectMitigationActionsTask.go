package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_startDetectMitigationActionsTaskCmd = &cobra.Command{
	Use:   "start-detect-mitigation-actions-task",
	Short: "Starts a Device Defender ML Detect mitigation actions task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_startDetectMitigationActionsTaskCmd).Standalone()

	iot_startDetectMitigationActionsTaskCmd.Flags().String("actions", "", "The actions to be performed when a device has unexpected behavior.")
	iot_startDetectMitigationActionsTaskCmd.Flags().String("client-request-token", "", "Each mitigation action task must have a unique client request token.")
	iot_startDetectMitigationActionsTaskCmd.Flags().String("include-only-active-violations", "", "Specifies to list only active violations.")
	iot_startDetectMitigationActionsTaskCmd.Flags().String("include-suppressed-alerts", "", "Specifies to include suppressed alerts.")
	iot_startDetectMitigationActionsTaskCmd.Flags().String("target", "", "Specifies the ML Detect findings to which the mitigation actions are applied.")
	iot_startDetectMitigationActionsTaskCmd.Flags().String("task-id", "", "The unique identifier of the task.")
	iot_startDetectMitigationActionsTaskCmd.Flags().String("violation-event-occurrence-range", "", "Specifies the time period of which violation events occurred between.")
	iot_startDetectMitigationActionsTaskCmd.MarkFlagRequired("actions")
	iot_startDetectMitigationActionsTaskCmd.MarkFlagRequired("client-request-token")
	iot_startDetectMitigationActionsTaskCmd.MarkFlagRequired("target")
	iot_startDetectMitigationActionsTaskCmd.MarkFlagRequired("task-id")
	iotCmd.AddCommand(iot_startDetectMitigationActionsTaskCmd)
}
