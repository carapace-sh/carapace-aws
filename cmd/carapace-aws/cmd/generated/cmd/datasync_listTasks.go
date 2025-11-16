package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_listTasksCmd = &cobra.Command{
	Use:   "list-tasks",
	Short: "Returns a list of the DataSync tasks you created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_listTasksCmd).Standalone()

	datasync_listTasksCmd.Flags().String("filters", "", "You can use API filters to narrow down the list of resources returned by `ListTasks`.")
	datasync_listTasksCmd.Flags().String("max-results", "", "The maximum number of tasks to return.")
	datasync_listTasksCmd.Flags().String("next-token", "", "An opaque string that indicates the position at which to begin the next list of tasks.")
	datasyncCmd.AddCommand(datasync_listTasksCmd)
}
