package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getExportTaskCmd = &cobra.Command{
	Use:   "get-export-task",
	Short: "Retrieves a specified export task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getExportTaskCmd).Standalone()

	neptuneGraph_getExportTaskCmd.Flags().String("task-identifier", "", "The unique identifier of the export task.")
	neptuneGraph_getExportTaskCmd.MarkFlagRequired("task-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_getExportTaskCmd)
}
