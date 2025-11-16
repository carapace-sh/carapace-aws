package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listThingRegistrationTasksCmd = &cobra.Command{
	Use:   "list-thing-registration-tasks",
	Short: "List bulk thing provisioning tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listThingRegistrationTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listThingRegistrationTasksCmd).Standalone()

		iot_listThingRegistrationTasksCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listThingRegistrationTasksCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listThingRegistrationTasksCmd.Flags().String("status", "", "The status of the bulk thing provisioning task.")
	})
	iotCmd.AddCommand(iot_listThingRegistrationTasksCmd)
}
