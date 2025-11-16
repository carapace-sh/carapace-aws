package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_stopClockCmd = &cobra.Command{
	Use:   "stop-clock",
	Short: "Stops the simulation clock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_stopClockCmd).Standalone()

	simspaceweaver_stopClockCmd.Flags().String("simulation", "", "The name of the simulation.")
	simspaceweaver_stopClockCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_stopClockCmd)
}
