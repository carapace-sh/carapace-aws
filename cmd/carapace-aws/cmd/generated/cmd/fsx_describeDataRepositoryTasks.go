package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeDataRepositoryTasksCmd = &cobra.Command{
	Use:   "describe-data-repository-tasks",
	Short: "Returns the description of specific Amazon FSx for Lustre or Amazon File Cache data repository tasks, if one or more `TaskIds` values are provided in the request, or if filters are used in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeDataRepositoryTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_describeDataRepositoryTasksCmd).Standalone()

		fsx_describeDataRepositoryTasksCmd.Flags().String("filters", "", "(Optional) You can use filters to narrow the `DescribeDataRepositoryTasks` response to include just tasks for specific file systems, or tasks in a specific lifecycle state.")
		fsx_describeDataRepositoryTasksCmd.Flags().String("max-results", "", "")
		fsx_describeDataRepositoryTasksCmd.Flags().String("next-token", "", "")
		fsx_describeDataRepositoryTasksCmd.Flags().String("task-ids", "", "(Optional) IDs of the tasks whose descriptions you want to retrieve (String).")
	})
	fsxCmd.AddCommand(fsx_describeDataRepositoryTasksCmd)
}
