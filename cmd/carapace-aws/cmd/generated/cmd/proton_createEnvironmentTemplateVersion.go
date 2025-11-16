package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createEnvironmentTemplateVersionCmd = &cobra.Command{
	Use:   "create-environment-template-version",
	Short: "Create a new major or minor version of an environment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createEnvironmentTemplateVersionCmd).Standalone()

	proton_createEnvironmentTemplateVersionCmd.Flags().String("client-token", "", "When included, if two identical requests are made with the same client token, Proton returns the environment template version that the first request created.")
	proton_createEnvironmentTemplateVersionCmd.Flags().String("description", "", "A description of the new version of an environment template.")
	proton_createEnvironmentTemplateVersionCmd.Flags().String("major-version", "", "To create a new minor version of the environment template, include `major Version`.")
	proton_createEnvironmentTemplateVersionCmd.Flags().String("source", "", "An object that includes the template bundle S3 bucket path and name for the new version of an template.")
	proton_createEnvironmentTemplateVersionCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton environment template version.")
	proton_createEnvironmentTemplateVersionCmd.Flags().String("template-name", "", "The name of the environment template.")
	proton_createEnvironmentTemplateVersionCmd.MarkFlagRequired("source")
	proton_createEnvironmentTemplateVersionCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_createEnvironmentTemplateVersionCmd)
}
