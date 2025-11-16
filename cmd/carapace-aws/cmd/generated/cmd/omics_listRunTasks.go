package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listRunTasksCmd = &cobra.Command{
	Use:   "list-run-tasks",
	Short: "Returns a list of tasks and status information within their specified run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listRunTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listRunTasksCmd).Standalone()

		omics_listRunTasksCmd.Flags().String("id", "", "The run's ID.")
		omics_listRunTasksCmd.Flags().String("max-results", "", "The maximum number of run tasks to return in one page of results.")
		omics_listRunTasksCmd.Flags().String("starting-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listRunTasksCmd.Flags().String("status", "", "Filter the list by status.")
		omics_listRunTasksCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_listRunTasksCmd)
}
