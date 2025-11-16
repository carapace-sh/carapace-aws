package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceTemplateVersionCmd = &cobra.Command{
	Use:   "get-service-template-version",
	Short: "Get detailed data for a major or minor version of a service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceTemplateVersionCmd).Standalone()

	proton_getServiceTemplateVersionCmd.Flags().String("major-version", "", "To get service template major version detail data, include `major Version`.")
	proton_getServiceTemplateVersionCmd.Flags().String("minor-version", "", "To get service template minor version detail data, include `minorVersion`.")
	proton_getServiceTemplateVersionCmd.Flags().String("template-name", "", "The name of the service template a version of which you want to get detailed data for.")
	proton_getServiceTemplateVersionCmd.MarkFlagRequired("major-version")
	proton_getServiceTemplateVersionCmd.MarkFlagRequired("minor-version")
	proton_getServiceTemplateVersionCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_getServiceTemplateVersionCmd)
}
