package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServiceTemplateVersionCmd = &cobra.Command{
	Use:   "update-service-template-version",
	Short: "Update a major or minor version of a service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServiceTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateServiceTemplateVersionCmd).Standalone()

		proton_updateServiceTemplateVersionCmd.Flags().String("compatible-environment-templates", "", "An array of environment template objects that are compatible with this service template version.")
		proton_updateServiceTemplateVersionCmd.Flags().String("description", "", "A description of a service template version to update.")
		proton_updateServiceTemplateVersionCmd.Flags().String("major-version", "", "To update a major version of a service template, include `major Version`.")
		proton_updateServiceTemplateVersionCmd.Flags().String("minor-version", "", "To update a minor version of a service template, include `minorVersion`.")
		proton_updateServiceTemplateVersionCmd.Flags().String("status", "", "The status of the service template minor version to update.")
		proton_updateServiceTemplateVersionCmd.Flags().String("supported-component-sources", "", "An array of supported component sources.")
		proton_updateServiceTemplateVersionCmd.Flags().String("template-name", "", "The name of the service template.")
		proton_updateServiceTemplateVersionCmd.MarkFlagRequired("major-version")
		proton_updateServiceTemplateVersionCmd.MarkFlagRequired("minor-version")
		proton_updateServiceTemplateVersionCmd.MarkFlagRequired("template-name")
	})
	protonCmd.AddCommand(proton_updateServiceTemplateVersionCmd)
}
