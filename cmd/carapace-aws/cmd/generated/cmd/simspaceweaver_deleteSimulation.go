package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_deleteSimulationCmd = &cobra.Command{
	Use:   "delete-simulation",
	Short: "Deletes all SimSpace Weaver resources assigned to the given simulation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_deleteSimulationCmd).Standalone()

	simspaceweaver_deleteSimulationCmd.Flags().String("simulation", "", "The name of the simulation.")
	simspaceweaver_deleteSimulationCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_deleteSimulationCmd)
}
