package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listEnvironmentTemplateVersionsCmd = &cobra.Command{
	Use:   "list-environment-template-versions",
	Short: "List major or minor versions of an environment template with detail data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listEnvironmentTemplateVersionsCmd).Standalone()

	proton_listEnvironmentTemplateVersionsCmd.Flags().String("major-version", "", "To view a list of minor of versions under a major version of an environment template, include `major Version`.")
	proton_listEnvironmentTemplateVersionsCmd.Flags().String("max-results", "", "The maximum number of major or minor versions of an environment template to list.")
	proton_listEnvironmentTemplateVersionsCmd.Flags().String("next-token", "", "A token that indicates the location of the next major or minor version in the array of major or minor versions of an environment template, after the list of major or minor versions that was previously requested.")
	proton_listEnvironmentTemplateVersionsCmd.Flags().String("template-name", "", "The name of the environment template.")
	proton_listEnvironmentTemplateVersionsCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_listEnvironmentTemplateVersionsCmd)
}
