package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_startClockCmd = &cobra.Command{
	Use:   "start-clock",
	Short: "Starts the simulation clock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_startClockCmd).Standalone()

	simspaceweaver_startClockCmd.Flags().String("simulation", "", "The name of the simulation.")
	simspaceweaver_startClockCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_startClockCmd)
}
