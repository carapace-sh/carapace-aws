package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Delete an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteEnvironmentCmd).Standalone()

		proton_deleteEnvironmentCmd.Flags().String("name", "", "The name of the environment to delete.")
		proton_deleteEnvironmentCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_deleteEnvironmentCmd)
}
