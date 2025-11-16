package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_listTaskExecutionsCmd = &cobra.Command{
	Use:   "list-task-executions",
	Short: "Returns a list of executions for an DataSync transfer task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_listTaskExecutionsCmd).Standalone()

	datasync_listTaskExecutionsCmd.Flags().String("max-results", "", "Specifies how many results you want in the response.")
	datasync_listTaskExecutionsCmd.Flags().String("next-token", "", "Specifies an opaque string that indicates the position at which to begin the next list of results in the response.")
	datasync_listTaskExecutionsCmd.Flags().String("task-arn", "", "Specifies the Amazon Resource Name (ARN) of the task that you want execution information about.")
	datasyncCmd.AddCommand(datasync_listTaskExecutionsCmd)
}
