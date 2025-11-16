package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listReceivedDataGrantsCmd = &cobra.Command{
	Use:   "list-received-data-grants",
	Short: "This operation returns information about all received data grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listReceivedDataGrantsCmd).Standalone()

	dataexchange_listReceivedDataGrantsCmd.Flags().String("acceptance-state", "", "The acceptance state of the data grants to list.")
	dataexchange_listReceivedDataGrantsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
	dataexchange_listReceivedDataGrantsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	dataexchangeCmd.AddCommand(dataexchange_listReceivedDataGrantsCmd)
}
