package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Deploy a new environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createEnvironmentCmd).Standalone()

	proton_createEnvironmentCmd.Flags().String("codebuild-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that allows Proton to provision infrastructure using CodeBuild-based provisioning on your behalf.")
	proton_createEnvironmentCmd.Flags().String("component-role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that Proton uses when provisioning directly defined components in this environment.")
	proton_createEnvironmentCmd.Flags().String("description", "", "A description of the environment that's being created and deployed.")
	proton_createEnvironmentCmd.Flags().String("environment-account-connection-id", "", "The ID of the environment account connection that you provide if you're provisioning your environment infrastructure resources to an environment account.")
	proton_createEnvironmentCmd.Flags().String("name", "", "The name of the environment.")
	proton_createEnvironmentCmd.Flags().String("proton-service-role-arn", "", "The Amazon Resource Name (ARN) of the Proton service role that allows Proton to make calls to other services on your behalf.")
	proton_createEnvironmentCmd.Flags().String("provisioning-repository", "", "The linked repository that you use to host your rendered infrastructure templates for self-managed provisioning.")
	proton_createEnvironmentCmd.Flags().String("spec", "", "A YAML formatted string that provides inputs as defined in the environment template bundle schema file.")
	proton_createEnvironmentCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton environment.")
	proton_createEnvironmentCmd.Flags().String("template-major-version", "", "The major version of the environment template.")
	proton_createEnvironmentCmd.Flags().String("template-minor-version", "", "The minor version of the environment template.")
	proton_createEnvironmentCmd.Flags().String("template-name", "", "The name of the environment template.")
	proton_createEnvironmentCmd.MarkFlagRequired("name")
	proton_createEnvironmentCmd.MarkFlagRequired("spec")
	proton_createEnvironmentCmd.MarkFlagRequired("template-major-version")
	proton_createEnvironmentCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_createEnvironmentCmd)
}
