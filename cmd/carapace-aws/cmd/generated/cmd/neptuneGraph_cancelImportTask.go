package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_cancelImportTaskCmd = &cobra.Command{
	Use:   "cancel-import-task",
	Short: "Deletes the specified import task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_cancelImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_cancelImportTaskCmd).Standalone()

		neptuneGraph_cancelImportTaskCmd.Flags().String("task-identifier", "", "The unique identifier of the import task.")
		neptuneGraph_cancelImportTaskCmd.MarkFlagRequired("task-identifier")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_cancelImportTaskCmd)
}
