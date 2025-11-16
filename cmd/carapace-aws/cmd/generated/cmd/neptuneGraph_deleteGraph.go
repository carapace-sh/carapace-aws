package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_deleteGraphCmd = &cobra.Command{
	Use:   "delete-graph",
	Short: "Deletes the specified graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_deleteGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_deleteGraphCmd).Standalone()

		neptuneGraph_deleteGraphCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_deleteGraphCmd.Flags().Bool("no-skip-snapshot", false, "Determines whether a final graph snapshot is created before the graph is deleted.")
		neptuneGraph_deleteGraphCmd.Flags().Bool("skip-snapshot", false, "Determines whether a final graph snapshot is created before the graph is deleted.")
		neptuneGraph_deleteGraphCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_deleteGraphCmd.Flag("no-skip-snapshot").Hidden = true
		neptuneGraph_deleteGraphCmd.MarkFlagRequired("no-skip-snapshot")
		neptuneGraph_deleteGraphCmd.MarkFlagRequired("skip-snapshot")
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_deleteGraphCmd)
}
