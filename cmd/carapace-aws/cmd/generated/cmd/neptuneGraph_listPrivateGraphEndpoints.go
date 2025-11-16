package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_listPrivateGraphEndpointsCmd = &cobra.Command{
	Use:   "list-private-graph-endpoints",
	Short: "Lists private endpoints for a specified Neptune Analytics graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_listPrivateGraphEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_listPrivateGraphEndpointsCmd).Standalone()

		neptuneGraph_listPrivateGraphEndpointsCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_listPrivateGraphEndpointsCmd.Flags().String("max-results", "", "The total number of records to return in the command's output.")
		neptuneGraph_listPrivateGraphEndpointsCmd.Flags().String("next-token", "", "Pagination token used to paginate output.")
		neptuneGraph_listPrivateGraphEndpointsCmd.MarkFlagRequired("graph-identifier")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_listPrivateGraphEndpointsCmd)
}
