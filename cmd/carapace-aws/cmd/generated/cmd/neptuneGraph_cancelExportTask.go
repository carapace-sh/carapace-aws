package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_cancelExportTaskCmd = &cobra.Command{
	Use:   "cancel-export-task",
	Short: "Cancel the specified export task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_cancelExportTaskCmd).Standalone()

	neptuneGraph_cancelExportTaskCmd.Flags().String("task-identifier", "", "The unique identifier of the export task.")
	neptuneGraph_cancelExportTaskCmd.MarkFlagRequired("task-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_cancelExportTaskCmd)
}
