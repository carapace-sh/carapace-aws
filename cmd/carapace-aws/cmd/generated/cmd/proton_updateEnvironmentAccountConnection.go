package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateEnvironmentAccountConnectionCmd = &cobra.Command{
	Use:   "update-environment-account-connection",
	Short: "In an environment account, update an environment account connection to use a new IAM role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateEnvironmentAccountConnectionCmd).Standalone()

	proton_updateEnvironmentAccountConnectionCmd.Flags().String("codebuild-role-arn", "", "The Amazon Resource Name (ARN) of an IAM service role in the environment account.")
	proton_updateEnvironmentAccountConnectionCmd.Flags().String("component-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that Proton uses when provisioning directly defined components in the associated environment account.")
	proton_updateEnvironmentAccountConnectionCmd.Flags().String("id", "", "The ID of the environment account connection to update.")
	proton_updateEnvironmentAccountConnectionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that's associated with the environment account connection to update.")
	proton_updateEnvironmentAccountConnectionCmd.MarkFlagRequired("id")
	protonCmd.AddCommand(proton_updateEnvironmentAccountConnectionCmd)
}
