package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getDataSetImportTaskCmd = &cobra.Command{
	Use:   "get-data-set-import-task",
	Short: "Gets the status of a data set import task initiated with the [CreateDataSetImportTask]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getDataSetImportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_getDataSetImportTaskCmd).Standalone()

		m2_getDataSetImportTaskCmd.Flags().String("application-id", "", "The application identifier.")
		m2_getDataSetImportTaskCmd.Flags().String("task-id", "", "The task identifier returned by the [CreateDataSetImportTask]() operation.")
		m2_getDataSetImportTaskCmd.MarkFlagRequired("application-id")
		m2_getDataSetImportTaskCmd.MarkFlagRequired("task-id")
	})
	m2Cmd.AddCommand(m2_getDataSetImportTaskCmd)
}
