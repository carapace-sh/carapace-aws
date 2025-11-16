package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServiceTemplateCmd = &cobra.Command{
	Use:   "update-service-template",
	Short: "Update a service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServiceTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateServiceTemplateCmd).Standalone()

		proton_updateServiceTemplateCmd.Flags().String("description", "", "A description of the service template update.")
		proton_updateServiceTemplateCmd.Flags().String("display-name", "", "The name of the service template to update that's displayed in the developer interface.")
		proton_updateServiceTemplateCmd.Flags().String("name", "", "The name of the service template to update.")
		proton_updateServiceTemplateCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_updateServiceTemplateCmd)
}
