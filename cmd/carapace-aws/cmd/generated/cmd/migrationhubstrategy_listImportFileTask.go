package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_listImportFileTaskCmd = &cobra.Command{
	Use:   "list-import-file-task",
	Short: "Retrieves a list of all the imports performed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_listImportFileTaskCmd).Standalone()

	migrationhubstrategy_listImportFileTaskCmd.Flags().String("max-results", "", "The total number of items to return.")
	migrationhubstrategy_listImportFileTaskCmd.Flags().String("next-token", "", "The token from a previous call that you use to retrieve the next set of results.")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_listImportFileTaskCmd)
}
