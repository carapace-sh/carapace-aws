package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_describeAppCmd = &cobra.Command{
	Use:   "describe-app",
	Short: "Returns the state of the given custom app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_describeAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(simspaceweaver_describeAppCmd).Standalone()

		simspaceweaver_describeAppCmd.Flags().String("app", "", "The name of the app.")
		simspaceweaver_describeAppCmd.Flags().String("domain", "", "The name of the domain of the app.")
		simspaceweaver_describeAppCmd.Flags().String("simulation", "", "The name of the simulation of the app.")
		simspaceweaver_describeAppCmd.MarkFlagRequired("app")
		simspaceweaver_describeAppCmd.MarkFlagRequired("domain")
		simspaceweaver_describeAppCmd.MarkFlagRequired("simulation")
	})
	simspaceweaverCmd.AddCommand(simspaceweaver_describeAppCmd)
}
