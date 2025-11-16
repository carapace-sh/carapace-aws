package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listImportTasksCmd = &cobra.Command{
	Use:   "list-import-tasks",
	Short: "Lists import tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listImportTasksCmd).Standalone()

	neptuneGraph_listImportTasksCmd.Flags().String("max-results", "", "The total number of records to return in the command's output.")
	neptuneGraph_listImportTasksCmd.Flags().String("next-token", "", "Pagination token used to paginate output.")
	neptuneGraphCmd.AddCommand(neptuneGraph_listImportTasksCmd)
}
