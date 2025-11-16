package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_describeSimulationCmd = &cobra.Command{
	Use:   "describe-simulation",
	Short: "Returns the current state of the given simulation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_describeSimulationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(simspaceweaver_describeSimulationCmd).Standalone()

		simspaceweaver_describeSimulationCmd.Flags().String("simulation", "", "The name of the simulation.")
		simspaceweaver_describeSimulationCmd.MarkFlagRequired("simulation")
	})
	simspaceweaverCmd.AddCommand(simspaceweaver_describeSimulationCmd)
}
