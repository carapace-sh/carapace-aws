package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_createDataSetImportTaskCmd = &cobra.Command{
	Use:   "create-data-set-import-task",
	Short: "Starts a data set import task for a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_createDataSetImportTaskCmd).Standalone()

	m2_createDataSetImportTaskCmd.Flags().String("application-id", "", "The unique identifier of the application for which you want to import data sets.")
	m2_createDataSetImportTaskCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request to create a data set import.")
	m2_createDataSetImportTaskCmd.Flags().String("import-config", "", "The data set import task configuration.")
	m2_createDataSetImportTaskCmd.MarkFlagRequired("application-id")
	m2_createDataSetImportTaskCmd.MarkFlagRequired("import-config")
	m2Cmd.AddCommand(m2_createDataSetImportTaskCmd)
}
