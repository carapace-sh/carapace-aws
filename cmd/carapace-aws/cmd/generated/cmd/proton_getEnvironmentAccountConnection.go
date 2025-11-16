package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getEnvironmentAccountConnectionCmd = &cobra.Command{
	Use:   "get-environment-account-connection",
	Short: "In an environment account, get the detailed data for an environment account connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getEnvironmentAccountConnectionCmd).Standalone()

	proton_getEnvironmentAccountConnectionCmd.Flags().String("id", "", "The ID of the environment account connection that you want to get the detailed data for.")
	proton_getEnvironmentAccountConnectionCmd.MarkFlagRequired("id")
	protonCmd.AddCommand(proton_getEnvironmentAccountConnectionCmd)
}
