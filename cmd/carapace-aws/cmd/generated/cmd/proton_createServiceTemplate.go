package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createServiceTemplateCmd = &cobra.Command{
	Use:   "create-service-template",
	Short: "Create a service template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createServiceTemplateCmd).Standalone()

	proton_createServiceTemplateCmd.Flags().String("description", "", "A description of the service template.")
	proton_createServiceTemplateCmd.Flags().String("display-name", "", "The name of the service template as displayed in the developer interface.")
	proton_createServiceTemplateCmd.Flags().String("encryption-key", "", "A customer provided encryption key that's used to encrypt data.")
	proton_createServiceTemplateCmd.Flags().String("name", "", "The name of the service template.")
	proton_createServiceTemplateCmd.Flags().String("pipeline-provisioning", "", "By default, Proton provides a service pipeline for your service.")
	proton_createServiceTemplateCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton service template.")
	proton_createServiceTemplateCmd.MarkFlagRequired("name")
	protonCmd.AddCommand(proton_createServiceTemplateCmd)
}
