package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Get detailed data for an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getEnvironmentCmd).Standalone()

		proton_getEnvironmentCmd.Flags().String("name", "", "The name of the environment that you want to get the detailed data for.")
		proton_getEnvironmentCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_getEnvironmentCmd)
}
