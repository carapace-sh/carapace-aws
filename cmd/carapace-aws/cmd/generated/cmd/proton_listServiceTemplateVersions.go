package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServiceTemplateVersionsCmd = &cobra.Command{
	Use:   "list-service-template-versions",
	Short: "List major or minor versions of a service template with detail data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServiceTemplateVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServiceTemplateVersionsCmd).Standalone()

		proton_listServiceTemplateVersionsCmd.Flags().String("major-version", "", "To view a list of minor of versions under a major version of a service template, include `major Version`.")
		proton_listServiceTemplateVersionsCmd.Flags().String("max-results", "", "The maximum number of major or minor versions of a service template to list.")
		proton_listServiceTemplateVersionsCmd.Flags().String("next-token", "", "A token that indicates the location of the next major or minor version in the array of major or minor versions of a service template, after the list of major or minor versions that was previously requested.")
		proton_listServiceTemplateVersionsCmd.Flags().String("template-name", "", "The name of the service template.")
		proton_listServiceTemplateVersionsCmd.MarkFlagRequired("template-name")
	})
	protonCmd.AddCommand(proton_listServiceTemplateVersionsCmd)
}
