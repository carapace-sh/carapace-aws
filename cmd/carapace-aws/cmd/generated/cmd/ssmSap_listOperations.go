package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listOperationsCmd = &cobra.Command{
	Use:   "list-operations",
	Short: "Lists the operations performed by AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_listOperationsCmd).Standalone()

		ssmSap_listOperationsCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_listOperationsCmd.Flags().String("filters", "", "The filters of an operation.")
		ssmSap_listOperationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ssmSap_listOperationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ssmSap_listOperationsCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_listOperationsCmd)
}
