package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingRegistrationTaskReportsCmd = &cobra.Command{
	Use:   "list-thing-registration-task-reports",
	Short: "Information about the thing registration tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingRegistrationTaskReportsCmd).Standalone()

	iot_listThingRegistrationTaskReportsCmd.Flags().String("max-results", "", "The maximum number of results to return per request.")
	iot_listThingRegistrationTaskReportsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iot_listThingRegistrationTaskReportsCmd.Flags().String("report-type", "", "The type of task report.")
	iot_listThingRegistrationTaskReportsCmd.Flags().String("task-id", "", "The id of the task.")
	iot_listThingRegistrationTaskReportsCmd.MarkFlagRequired("report-type")
	iot_listThingRegistrationTaskReportsCmd.MarkFlagRequired("task-id")
	iotCmd.AddCommand(iot_listThingRegistrationTaskReportsCmd)
}
