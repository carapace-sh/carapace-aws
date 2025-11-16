package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteEnvironmentTemplateVersionCmd = &cobra.Command{
	Use:   "delete-environment-template-version",
	Short: "If no other minor versions of an environment template exist, delete a major version of the environment template if it's not the `Recommended` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteEnvironmentTemplateVersionCmd).Standalone()

	proton_deleteEnvironmentTemplateVersionCmd.Flags().String("major-version", "", "The environment template major version to delete.")
	proton_deleteEnvironmentTemplateVersionCmd.Flags().String("minor-version", "", "The environment template minor version to delete.")
	proton_deleteEnvironmentTemplateVersionCmd.Flags().String("template-name", "", "The name of the environment template.")
	proton_deleteEnvironmentTemplateVersionCmd.MarkFlagRequired("major-version")
	proton_deleteEnvironmentTemplateVersionCmd.MarkFlagRequired("minor-version")
	proton_deleteEnvironmentTemplateVersionCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_deleteEnvironmentTemplateVersionCmd)
}
