package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listDetectMitigationActionsExecutionsCmd = &cobra.Command{
	Use:   "list-detect-mitigation-actions-executions",
	Short: "Lists mitigation actions executions for a Device Defender ML Detect Security Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listDetectMitigationActionsExecutionsCmd).Standalone()

	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("end-time", "", "The end of the time period for which ML Detect mitigation actions executions are returned.")
	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("start-time", "", "A filter to limit results to those found after the specified time.")
	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("task-id", "", "The unique identifier of the task.")
	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("thing-name", "", "The name of the thing whose mitigation actions are listed.")
	iot_listDetectMitigationActionsExecutionsCmd.Flags().String("violation-id", "", "The unique identifier of the violation.")
	iotCmd.AddCommand(iot_listDetectMitigationActionsExecutionsCmd)
}
