package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listOperationEventsCmd = &cobra.Command{
	Use:   "list-operation-events",
	Short: "Returns a list of operations events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listOperationEventsCmd).Standalone()

	ssmSap_listOperationEventsCmd.Flags().String("filters", "", "Optionally specify filters to narrow the returned operation event items.")
	ssmSap_listOperationEventsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ssmSap_listOperationEventsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	ssmSap_listOperationEventsCmd.Flags().String("operation-id", "", "The ID of the operation.")
	ssmSap_listOperationEventsCmd.MarkFlagRequired("operation-id")
	ssmSapCmd.AddCommand(ssmSap_listOperationEventsCmd)
}
