package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_getGraphCmd = &cobra.Command{
	Use:   "get-graph",
	Short: "Gets information about a specified graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_getGraphCmd).Standalone()

	neptuneGraph_getGraphCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_getGraphCmd.MarkFlagRequired("graph-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_getGraphCmd)
}
