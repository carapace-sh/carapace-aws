package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listDataSetImportHistoryCmd = &cobra.Command{
	Use:   "list-data-set-import-history",
	Short: "Lists the data set imports for the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listDataSetImportHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_listDataSetImportHistoryCmd).Standalone()

		m2_listDataSetImportHistoryCmd.Flags().String("application-id", "", "The unique identifier of the application.")
		m2_listDataSetImportHistoryCmd.Flags().String("max-results", "", "The maximum number of objects to return.")
		m2_listDataSetImportHistoryCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
		m2_listDataSetImportHistoryCmd.MarkFlagRequired("application-id")
	})
	m2Cmd.AddCommand(m2_listDataSetImportHistoryCmd)
}
