package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_rejectEnvironmentAccountConnectionCmd = &cobra.Command{
	Use:   "reject-environment-account-connection",
	Short: "In a management account, reject an environment account connection from another environment account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_rejectEnvironmentAccountConnectionCmd).Standalone()

	proton_rejectEnvironmentAccountConnectionCmd.Flags().String("id", "", "The ID of the environment account connection to reject.")
	proton_rejectEnvironmentAccountConnectionCmd.MarkFlagRequired("id")
	protonCmd.AddCommand(proton_rejectEnvironmentAccountConnectionCmd)
}
