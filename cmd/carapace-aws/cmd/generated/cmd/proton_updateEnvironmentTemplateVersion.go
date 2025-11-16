package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateEnvironmentTemplateVersionCmd = &cobra.Command{
	Use:   "update-environment-template-version",
	Short: "Update a major or minor version of an environment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateEnvironmentTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateEnvironmentTemplateVersionCmd).Standalone()

		proton_updateEnvironmentTemplateVersionCmd.Flags().String("description", "", "A description of environment template version to update.")
		proton_updateEnvironmentTemplateVersionCmd.Flags().String("major-version", "", "To update a major version of an environment template, include `major Version`.")
		proton_updateEnvironmentTemplateVersionCmd.Flags().String("minor-version", "", "To update a minor version of an environment template, include `minorVersion`.")
		proton_updateEnvironmentTemplateVersionCmd.Flags().String("status", "", "The status of the environment template minor version to update.")
		proton_updateEnvironmentTemplateVersionCmd.Flags().String("template-name", "", "The name of the environment template.")
		proton_updateEnvironmentTemplateVersionCmd.MarkFlagRequired("major-version")
		proton_updateEnvironmentTemplateVersionCmd.MarkFlagRequired("minor-version")
		proton_updateEnvironmentTemplateVersionCmd.MarkFlagRequired("template-name")
	})
	protonCmd.AddCommand(proton_updateEnvironmentTemplateVersionCmd)
}
