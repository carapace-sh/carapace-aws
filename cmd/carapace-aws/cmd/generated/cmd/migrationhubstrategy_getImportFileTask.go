package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getImportFileTaskCmd = &cobra.Command{
	Use:   "get-import-file-task",
	Short: "Retrieves the details about a specific import task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getImportFileTaskCmd).Standalone()

	migrationhubstrategy_getImportFileTaskCmd.Flags().String("id", "", "The ID of the import file task.")
	migrationhubstrategy_getImportFileTaskCmd.MarkFlagRequired("id")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getImportFileTaskCmd)
}
