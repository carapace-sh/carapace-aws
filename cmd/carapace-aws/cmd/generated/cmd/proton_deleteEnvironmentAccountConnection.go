package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteEnvironmentAccountConnectionCmd = &cobra.Command{
	Use:   "delete-environment-account-connection",
	Short: "In an environment account, delete an environment account connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteEnvironmentAccountConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteEnvironmentAccountConnectionCmd).Standalone()

		proton_deleteEnvironmentAccountConnectionCmd.Flags().String("id", "", "The ID of the environment account connection to delete.")
		proton_deleteEnvironmentAccountConnectionCmd.MarkFlagRequired("id")
	})
	protonCmd.AddCommand(proton_deleteEnvironmentAccountConnectionCmd)
}
