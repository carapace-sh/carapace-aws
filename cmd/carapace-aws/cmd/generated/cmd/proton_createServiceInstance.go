package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createServiceInstanceCmd = &cobra.Command{
	Use:   "create-service-instance",
	Short: "Create a service instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createServiceInstanceCmd).Standalone()

	proton_createServiceInstanceCmd.Flags().String("client-token", "", "The client token of the service instance to create.")
	proton_createServiceInstanceCmd.Flags().String("name", "", "The name of the service instance to create.")
	proton_createServiceInstanceCmd.Flags().String("service-name", "", "The name of the service the service instance is added to.")
	proton_createServiceInstanceCmd.Flags().String("spec", "", "The spec for the service instance you want to create.")
	proton_createServiceInstanceCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton service instance.")
	proton_createServiceInstanceCmd.Flags().String("template-major-version", "", "To create a new major and minor version of the service template, *exclude* `major Version`.")
	proton_createServiceInstanceCmd.Flags().String("template-minor-version", "", "To create a new minor version of the service template, include a `major Version`.")
	proton_createServiceInstanceCmd.MarkFlagRequired("name")
	proton_createServiceInstanceCmd.MarkFlagRequired("service-name")
	proton_createServiceInstanceCmd.MarkFlagRequired("spec")
	protonCmd.AddCommand(proton_createServiceInstanceCmd)
}
