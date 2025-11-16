package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listDataSetExportHistoryCmd = &cobra.Command{
	Use:   "list-data-set-export-history",
	Short: "Lists the data set exports for the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listDataSetExportHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_listDataSetExportHistoryCmd).Standalone()

		m2_listDataSetExportHistoryCmd.Flags().String("application-id", "", "The unique identifier of the application.")
		m2_listDataSetExportHistoryCmd.Flags().String("max-results", "", "The maximum number of objects to return.")
		m2_listDataSetExportHistoryCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
		m2_listDataSetExportHistoryCmd.MarkFlagRequired("application-id")
	})
	m2Cmd.AddCommand(m2_listDataSetExportHistoryCmd)
}
