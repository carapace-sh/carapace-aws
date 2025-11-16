package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listSubCheckResultsCmd = &cobra.Command{
	Use:   "list-sub-check-results",
	Short: "Lists the sub-check results of a specified configuration check operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listSubCheckResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_listSubCheckResultsCmd).Standalone()

		ssmSap_listSubCheckResultsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ssmSap_listSubCheckResultsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ssmSap_listSubCheckResultsCmd.Flags().String("operation-id", "", "The ID of the configuration check operation.")
		ssmSap_listSubCheckResultsCmd.MarkFlagRequired("operation-id")
	})
	ssmSapCmd.AddCommand(ssmSap_listSubCheckResultsCmd)
}
