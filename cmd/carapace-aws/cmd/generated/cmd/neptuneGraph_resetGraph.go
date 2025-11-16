package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_resetGraphCmd = &cobra.Command{
	Use:   "reset-graph",
	Short: "Empties the data from a specified Neptune Analytics graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_resetGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_resetGraphCmd).Standalone()

		neptuneGraph_resetGraphCmd.Flags().String("graph-identifier", "", "ID of the graph to reset.")
		neptuneGraph_resetGraphCmd.Flags().Bool("no-skip-snapshot", false, "Determines whether a final graph snapshot is created before the graph data is deleted.")
		neptuneGraph_resetGraphCmd.Flags().Bool("skip-snapshot", false, "Determines whether a final graph snapshot is created before the graph data is deleted.")
		neptuneGraph_resetGraphCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_resetGraphCmd.Flag("no-skip-snapshot").Hidden = true
		neptuneGraph_resetGraphCmd.MarkFlagRequired("no-skip-snapshot")
		neptuneGraph_resetGraphCmd.MarkFlagRequired("skip-snapshot")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_resetGraphCmd)
}
