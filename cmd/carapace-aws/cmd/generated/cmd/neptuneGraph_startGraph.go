package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_startGraphCmd = &cobra.Command{
	Use:   "start-graph",
	Short: "Starts the specific graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_startGraphCmd).Standalone()

	neptuneGraph_startGraphCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
	neptuneGraph_startGraphCmd.MarkFlagRequired("graph-identifier")
	neptuneGraphCmd.AddCommand(neptuneGraph_startGraphCmd)
}
