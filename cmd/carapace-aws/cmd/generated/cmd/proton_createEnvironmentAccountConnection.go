package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createEnvironmentAccountConnectionCmd = &cobra.Command{
	Use:   "create-environment-account-connection",
	Short: "Create an environment account connection in an environment account so that environment infrastructure resources can be provisioned in the environment account from a management account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createEnvironmentAccountConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_createEnvironmentAccountConnectionCmd).Standalone()

		proton_createEnvironmentAccountConnectionCmd.Flags().String("client-token", "", "When included, if two identical requests are made with the same client token, Proton returns the environment account connection that the first request created.")
		proton_createEnvironmentAccountConnectionCmd.Flags().String("codebuild-role-arn", "", "The Amazon Resource Name (ARN) of an IAM service role in the environment account.")
		proton_createEnvironmentAccountConnectionCmd.Flags().String("component-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that Proton uses when provisioning directly defined components in the associated environment account.")
		proton_createEnvironmentAccountConnectionCmd.Flags().String("environment-name", "", "The name of the Proton environment that's created in the associated management account.")
		proton_createEnvironmentAccountConnectionCmd.Flags().String("management-account-id", "", "The ID of the management account that accepts or rejects the environment account connection.")
		proton_createEnvironmentAccountConnectionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account.")
		proton_createEnvironmentAccountConnectionCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton environment account connection.")
		proton_createEnvironmentAccountConnectionCmd.MarkFlagRequired("environment-name")
		proton_createEnvironmentAccountConnectionCmd.MarkFlagRequired("management-account-id")
	})
	protonCmd.AddCommand(proton_createEnvironmentAccountConnectionCmd)
}
