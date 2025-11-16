package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getDataSetExportTaskCmd = &cobra.Command{
	Use:   "get-data-set-export-task",
	Short: "Gets the status of a data set import task initiated with the [CreateDataSetExportTask]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getDataSetExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_getDataSetExportTaskCmd).Standalone()

		m2_getDataSetExportTaskCmd.Flags().String("application-id", "", "The application identifier.")
		m2_getDataSetExportTaskCmd.Flags().String("task-id", "", "The task identifier returned by the [CreateDataSetExportTask]() operation.")
		m2_getDataSetExportTaskCmd.MarkFlagRequired("application-id")
		m2_getDataSetExportTaskCmd.MarkFlagRequired("task-id")
	})
	m2Cmd.AddCommand(m2_getDataSetExportTaskCmd)
}
