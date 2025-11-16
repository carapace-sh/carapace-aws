package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_stopGraphCmd = &cobra.Command{
	Use:   "stop-graph",
	Short: "Stops the specific graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_stopGraphCmd).Standalone()

	neptuneGraph_stopGraphCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_stopGraphCmd.MarkFlagRequired("graph-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_stopGraphCmd)
}
