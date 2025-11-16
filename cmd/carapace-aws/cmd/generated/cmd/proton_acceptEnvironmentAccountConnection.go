package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_acceptEnvironmentAccountConnectionCmd = &cobra.Command{
	Use:   "accept-environment-account-connection",
	Short: "In a management account, an environment account connection request is accepted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_acceptEnvironmentAccountConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_acceptEnvironmentAccountConnectionCmd).Standalone()

		proton_acceptEnvironmentAccountConnectionCmd.Flags().String("id", "", "The ID of the environment account connection.")
		proton_acceptEnvironmentAccountConnectionCmd.MarkFlagRequired("id")
	})
	protonCmd.AddCommand(proton_acceptEnvironmentAccountConnectionCmd)
}
