package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Update an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateEnvironmentCmd).Standalone()

		proton_updateEnvironmentCmd.Flags().String("codebuild-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that allows Proton to provision infrastructure using CodeBuild-based provisioning on your behalf.")
		proton_updateEnvironmentCmd.Flags().String("component-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that Proton uses when provisioning directly defined components in this environment.")
		proton_updateEnvironmentCmd.Flags().String("deployment-type", "", "There are four modes for updating an environment.")
		proton_updateEnvironmentCmd.Flags().String("description", "", "A description of the environment update.")
		proton_updateEnvironmentCmd.Flags().String("environment-account-connection-id", "", "The ID of the environment account connection.")
		proton_updateEnvironmentCmd.Flags().String("name", "", "The name of the environment to update.")
		proton_updateEnvironmentCmd.Flags().String("proton-service-role-arn", "", "The Amazon Resource Name (ARN) of the Proton service role that allows Proton to make API calls to other services your behalf.")
		proton_updateEnvironmentCmd.Flags().String("provisioning-repository", "", "The linked repository that you use to host your rendered infrastructure templates for self-managed provisioning.")
		proton_updateEnvironmentCmd.Flags().String("spec", "", "The formatted specification that defines the update.")
		proton_updateEnvironmentCmd.Flags().String("template-major-version", "", "The major version of the environment to update.")
		proton_updateEnvironmentCmd.Flags().String("template-minor-version", "", "The minor version of the environment to update.")
		proton_updateEnvironmentCmd.MarkFlagRequired("deployment-type")
		proton_updateEnvironmentCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_updateEnvironmentCmd)
}
