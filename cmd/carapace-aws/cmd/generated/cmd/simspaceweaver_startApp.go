package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_startAppCmd = &cobra.Command{
	Use:   "start-app",
	Short: "Starts a custom app with the configuration specified in the simulation schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_startAppCmd).Standalone()

	simspaceweaver_startAppCmd.Flags().String("client-token", "", "A value that you provide to ensure that repeated calls to this API operation using the same parameters complete only once.")
	simspaceweaver_startAppCmd.Flags().String("description", "", "The description of the app.")
	simspaceweaver_startAppCmd.Flags().String("domain", "", "The name of the domain of the app.")
	simspaceweaver_startAppCmd.Flags().String("launch-overrides", "", "")
	simspaceweaver_startAppCmd.Flags().String("name", "", "The name of the app.")
	simspaceweaver_startAppCmd.Flags().String("simulation", "", "The name of the simulation of the app.")
	simspaceweaver_startAppCmd.MarkFlagRequired("domain")
	simspaceweaver_startAppCmd.MarkFlagRequired("name")
	simspaceweaver_startAppCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_startAppCmd)
}
