package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getEnvironmentTemplateVersionCmd = &cobra.Command{
	Use:   "get-environment-template-version",
	Short: "Get detailed data for a major or minor version of an environment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getEnvironmentTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getEnvironmentTemplateVersionCmd).Standalone()

		proton_getEnvironmentTemplateVersionCmd.Flags().String("major-version", "", "To get environment template major version detail data, include `major Version`.")
		proton_getEnvironmentTemplateVersionCmd.Flags().String("minor-version", "", "To get environment template minor version detail data, include `minorVersion`.")
		proton_getEnvironmentTemplateVersionCmd.Flags().String("template-name", "", "The name of the environment template a version of which you want to get detailed data for.")
		proton_getEnvironmentTemplateVersionCmd.MarkFlagRequired("major-version")
		proton_getEnvironmentTemplateVersionCmd.MarkFlagRequired("minor-version")
		proton_getEnvironmentTemplateVersionCmd.MarkFlagRequired("template-name")
	})
	protonCmd.AddCommand(proton_getEnvironmentTemplateVersionCmd)
}
