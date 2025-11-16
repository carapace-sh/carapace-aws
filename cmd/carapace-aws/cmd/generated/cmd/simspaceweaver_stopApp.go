package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_stopAppCmd = &cobra.Command{
	Use:   "stop-app",
	Short: "Stops the given custom app and shuts down all of its allocated compute resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_stopAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(simspaceweaver_stopAppCmd).Standalone()

		simspaceweaver_stopAppCmd.Flags().String("app", "", "The name of the app.")
		simspaceweaver_stopAppCmd.Flags().String("domain", "", "The name of the domain of the app.")
		simspaceweaver_stopAppCmd.Flags().String("simulation", "", "The name of the simulation of the app.")
		simspaceweaver_stopAppCmd.MarkFlagRequired("app")
		simspaceweaver_stopAppCmd.MarkFlagRequired("domain")
		simspaceweaver_stopAppCmd.MarkFlagRequired("simulation")
	})
	simspaceweaverCmd.AddCommand(simspaceweaver_stopAppCmd)
}
