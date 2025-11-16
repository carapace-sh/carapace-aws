package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listDataGrantsCmd = &cobra.Command{
	Use:   "list-data-grants",
	Short: "This operation returns information about all data grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listDataGrantsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_listDataGrantsCmd).Standalone()

		dataexchange_listDataGrantsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		dataexchange_listDataGrantsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	dataexchangeCmd.AddCommand(dataexchange_listDataGrantsCmd)
}
