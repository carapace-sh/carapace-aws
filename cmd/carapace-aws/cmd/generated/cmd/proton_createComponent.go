package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Create an Proton component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createComponentCmd).Standalone()

	proton_createComponentCmd.Flags().String("client-token", "", "The client token for the created component.")
	proton_createComponentCmd.Flags().String("description", "", "An optional customer-provided description of the component.")
	proton_createComponentCmd.Flags().String("environment-name", "", "The name of the Proton environment that you want to associate this component with.")
	proton_createComponentCmd.Flags().String("manifest", "", "A path to a manifest file that lists the Infrastructure as Code (IaC) file, template language, and rendering engine for infrastructure that a custom component provisions.")
	proton_createComponentCmd.Flags().String("name", "", "The customer-provided name of the component.")
	proton_createComponentCmd.Flags().String("service-instance-name", "", "The name of the service instance that you want to attach this component to.")
	proton_createComponentCmd.Flags().String("service-name", "", "The name of the service that `serviceInstanceName` is associated with.")
	proton_createComponentCmd.Flags().String("service-spec", "", "The service spec that you want the component to use to access service inputs.")
	proton_createComponentCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton component.")
	proton_createComponentCmd.Flags().String("template-file", "", "A path to the Infrastructure as Code (IaC) file describing infrastructure that a custom component provisions.")
	proton_createComponentCmd.MarkFlagRequired("manifest")
	proton_createComponentCmd.MarkFlagRequired("name")
	proton_createComponentCmd.MarkFlagRequired("template-file")
	protonCmd.AddCommand(proton_createComponentCmd)
}
