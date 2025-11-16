package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Deletes the instance of the given custom app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_deleteAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(simspaceweaver_deleteAppCmd).Standalone()

		simspaceweaver_deleteAppCmd.Flags().String("app", "", "The name of the app.")
		simspaceweaver_deleteAppCmd.Flags().String("domain", "", "The name of the domain of the app.")
		simspaceweaver_deleteAppCmd.Flags().String("simulation", "", "The name of the simulation of the app.")
		simspaceweaver_deleteAppCmd.MarkFlagRequired("app")
		simspaceweaver_deleteAppCmd.MarkFlagRequired("domain")
		simspaceweaver_deleteAppCmd.MarkFlagRequired("simulation")
	})
	simspaceweaverCmd.AddCommand(simspaceweaver_deleteAppCmd)
}
