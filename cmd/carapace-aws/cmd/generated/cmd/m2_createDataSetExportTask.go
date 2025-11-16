package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_createDataSetExportTaskCmd = &cobra.Command{
	Use:   "create-data-set-export-task",
	Short: "Starts a data set export task for a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_createDataSetExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_createDataSetExportTaskCmd).Standalone()

		m2_createDataSetExportTaskCmd.Flags().String("application-id", "", "The unique identifier of the application for which you want to export data sets.")
		m2_createDataSetExportTaskCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request to create a data set export.")
		m2_createDataSetExportTaskCmd.Flags().String("export-config", "", "The data set export task configuration.")
		m2_createDataSetExportTaskCmd.Flags().String("kms-key-id", "", "The identifier of a customer managed key.")
		m2_createDataSetExportTaskCmd.MarkFlagRequired("application-id")
		m2_createDataSetExportTaskCmd.MarkFlagRequired("export-config")
	})
	m2Cmd.AddCommand(m2_createDataSetExportTaskCmd)
}
