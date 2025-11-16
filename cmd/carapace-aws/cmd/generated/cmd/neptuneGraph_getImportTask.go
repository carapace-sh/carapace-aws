package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getImportTaskCmd = &cobra.Command{
	Use:   "get-import-task",
	Short: "Retrieves a specified import task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_getImportTaskCmd).Standalone()

		neptuneGraph_getImportTaskCmd.Flags().String("task-identifier", "", "The unique identifier of the import task.")
		neptuneGraph_getImportTaskCmd.MarkFlagRequired("task-identifier")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_getImportTaskCmd)
}
