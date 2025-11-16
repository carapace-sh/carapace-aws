package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_listSimulationsCmd = &cobra.Command{
	Use:   "list-simulations",
	Short: "Lists the SimSpace Weaver simulations in the Amazon Web Services account used to make the API call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_listSimulationsCmd).Standalone()

	simspaceweaver_listSimulationsCmd.Flags().String("max-results", "", "The maximum number of simulations to list.")
	simspaceweaver_listSimulationsCmd.Flags().String("next-token", "", "If SimSpace Weaver returns `nextToken`, then there are more results available.")
	simspaceweaverCmd.AddCommand(simspaceweaver_listSimulationsCmd)
}
