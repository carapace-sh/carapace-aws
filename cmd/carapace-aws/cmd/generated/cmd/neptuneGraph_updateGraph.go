package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_updateGraphCmd = &cobra.Command{
	Use:   "update-graph",
	Short: "Updates the configuration of a specified Neptune Analytics graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_updateGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraph_updateGraphCmd).Standalone()

		neptuneGraph_updateGraphCmd.Flags().Bool("deletion-protection", false, "A value that indicates whether the graph has deletion protection enabled.")
		neptuneGraph_updateGraphCmd.Flags().String("graph-identifier", "", "The unique identifier of the Neptune Analytics graph.")
		neptuneGraph_updateGraphCmd.Flags().Bool("no-deletion-protection", false, "A value that indicates whether the graph has deletion protection enabled.")
		neptuneGraph_updateGraphCmd.Flags().Bool("no-public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
		neptuneGraph_updateGraphCmd.Flags().String("provisioned-memory", "", "The provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.")
		neptuneGraph_updateGraphCmd.Flags().Bool("public-connectivity", false, "Specifies whether or not the graph can be reachable over the internet.")
		neptuneGraph_updateGraphCmd.MarkFlagRequired("graph-identifier")
		neptuneGraph_updateGraphCmd.Flag("no-deletion-protection").Hidden = true
		neptuneGraph_updateGraphCmd.Flag("no-public-connectivity").Hidden = true
	})
	neptuneGraphCmd.AddCommand(neptuneGraph_updateGraphCmd)
}
