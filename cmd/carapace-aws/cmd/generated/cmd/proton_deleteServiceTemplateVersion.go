package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteServiceTemplateVersionCmd = &cobra.Command{
	Use:   "delete-service-template-version",
	Short: "If no other minor versions of a service template exist, delete a major version of the service template if it's not the `Recommended` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteServiceTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteServiceTemplateVersionCmd).Standalone()

		proton_deleteServiceTemplateVersionCmd.Flags().String("major-version", "", "The service template major version to delete.")
		proton_deleteServiceTemplateVersionCmd.Flags().String("minor-version", "", "The service template minor version to delete.")
		proton_deleteServiceTemplateVersionCmd.Flags().String("template-name", "", "The name of the service template.")
		proton_deleteServiceTemplateVersionCmd.MarkFlagRequired("major-version")
		proton_deleteServiceTemplateVersionCmd.MarkFlagRequired("minor-version")
		proton_deleteServiceTemplateVersionCmd.MarkFlagRequired("template-name")
	})
	protonCmd.AddCommand(proton_deleteServiceTemplateVersionCmd)
}
