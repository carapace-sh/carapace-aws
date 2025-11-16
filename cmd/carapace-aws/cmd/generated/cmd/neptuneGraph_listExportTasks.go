package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listExportTasksCmd = &cobra.Command{
	Use:   "list-export-tasks",
	Short: "Retrieves a list of export tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listExportTasksCmd).Standalone()

	neptuneGraph_listExportTasksCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_listExportTasksCmd.Flags().String("max-results", "", "The maximum number of export tasks to return.")
	neptuneGraph_listExportTasksCmd.Flags().String("next-token", "", "Pagination token used to paginate input.")
	neptuneGraphCmd.AddCommand(neptuneGraph_listExportTasksCmd)
}
