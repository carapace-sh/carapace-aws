package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_listTasksCmd = &cobra.Command{
	Use:   "list-tasks",
	Short: "Returns a list of tasks that can be filtered by state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_listTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_listTasksCmd).Standalone()

		snowDeviceManagement_listTasksCmd.Flags().String("max-results", "", "The maximum number of tasks per page.")
		snowDeviceManagement_listTasksCmd.Flags().String("next-token", "", "A pagination token to continue to the next page of tasks.")
		snowDeviceManagement_listTasksCmd.Flags().String("state", "", "A structure used to filter the list of tasks.")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_listTasksCmd)
}
