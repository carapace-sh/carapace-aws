package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_cancelQueryCmd = &cobra.Command{
	Use:   "cancel-query",
	Short: "Cancels a specified query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_cancelQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_cancelQueryCmd).Standalone()

		neptuneGraph_cancelQueryCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_cancelQueryCmd.Flags().String("query-id", "", "The unique identifier of the query to cancel.")
		neptuneGraph_cancelQueryCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_cancelQueryCmd.MarkFlagRequired("query-id")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_cancelQueryCmd)
}
