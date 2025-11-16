package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listDetectMitigationActionsTasksCmd = &cobra.Command{
	Use:   "list-detect-mitigation-actions-tasks",
	Short: "List of Device Defender ML Detect mitigation actions tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listDetectMitigationActionsTasksCmd).Standalone()

	iot_listDetectMitigationActionsTasksCmd.Flags().String("end-time", "", "The end of the time period for which ML Detect mitigation actions tasks are returned.")
	iot_listDetectMitigationActionsTasksCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listDetectMitigationActionsTasksCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listDetectMitigationActionsTasksCmd.Flags().String("start-time", "", "A filter to limit results to those found after the specified time.")
	iot_listDetectMitigationActionsTasksCmd.MarkFlagRequired("end-time")
	iot_listDetectMitigationActionsTasksCmd.MarkFlagRequired("start-time")
	iotCmd.AddCommand(iot_listDetectMitigationActionsTasksCmd)
}
