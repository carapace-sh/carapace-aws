package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createServiceTemplateVersionCmd = &cobra.Command{
	Use:   "create-service-template-version",
	Short: "Create a new major or minor version of a service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createServiceTemplateVersionCmd).Standalone()

	proton_createServiceTemplateVersionCmd.Flags().String("client-token", "", "When included, if two identical requests are made with the same client token, Proton returns the service template version that the first request created.")
	proton_createServiceTemplateVersionCmd.Flags().String("compatible-environment-templates", "", "An array of environment template objects that are compatible with the new service template version.")
	proton_createServiceTemplateVersionCmd.Flags().String("description", "", "A description of the new version of a service template.")
	proton_createServiceTemplateVersionCmd.Flags().String("major-version", "", "To create a new minor version of the service template, include a `major Version`.")
	proton_createServiceTemplateVersionCmd.Flags().String("source", "", "An object that includes the template bundle S3 bucket path and name for the new version of a service template.")
	proton_createServiceTemplateVersionCmd.Flags().String("supported-component-sources", "", "An array of supported component sources.")
	proton_createServiceTemplateVersionCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton service template version.")
	proton_createServiceTemplateVersionCmd.Flags().String("template-name", "", "The name of the service template.")
	proton_createServiceTemplateVersionCmd.MarkFlagRequired("compatible-environment-templates")
	proton_createServiceTemplateVersionCmd.MarkFlagRequired("source")
	proton_createServiceTemplateVersionCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_createServiceTemplateVersionCmd)
}
