package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_stopSimulationCmd = &cobra.Command{
	Use:   "stop-simulation",
	Short: "Stops the given simulation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_stopSimulationCmd).Standalone()

	simspaceweaver_stopSimulationCmd.Flags().String("simulation", "", "The name of the simulation.")
	simspaceweaver_stopSimulationCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_stopSimulationCmd)
}
