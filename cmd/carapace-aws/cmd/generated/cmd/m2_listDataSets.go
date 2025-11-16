package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listDataSetsCmd = &cobra.Command{
	Use:   "list-data-sets",
	Short: "Lists the data sets imported for a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listDataSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_listDataSetsCmd).Standalone()

		m2_listDataSetsCmd.Flags().String("application-id", "", "The unique identifier of the application for which you want to list the associated data sets.")
		m2_listDataSetsCmd.Flags().String("max-results", "", "The maximum number of objects to return.")
		m2_listDataSetsCmd.Flags().String("name-filter", "", "Filter dataset name matching the specified pattern.")
		m2_listDataSetsCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
		m2_listDataSetsCmd.Flags().String("prefix", "", "The prefix of the data set name, which you can use to filter the list of data sets.")
		m2_listDataSetsCmd.MarkFlagRequired("application-id")
	})
	m2Cmd.AddCommand(m2_listDataSetsCmd)
}
