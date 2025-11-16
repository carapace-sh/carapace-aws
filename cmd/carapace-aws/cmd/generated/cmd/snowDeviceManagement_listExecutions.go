package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_listExecutionsCmd = &cobra.Command{
	Use:   "list-executions",
	Short: "Returns the status of tasks for one or more target devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_listExecutionsCmd).Standalone()

	snowDeviceManagement_listExecutionsCmd.Flags().String("max-results", "", "The maximum number of tasks to list per page.")
	snowDeviceManagement_listExecutionsCmd.Flags().String("next-token", "", "A pagination token to continue to the next page of tasks.")
	snowDeviceManagement_listExecutionsCmd.Flags().String("state", "", "A structure used to filter the tasks by their current state.")
	snowDeviceManagement_listExecutionsCmd.Flags().String("task-id", "", "The ID of the task.")
	snowDeviceManagement_listExecutionsCmd.MarkFlagRequired("task-id")
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_listExecutionsCmd)
}
