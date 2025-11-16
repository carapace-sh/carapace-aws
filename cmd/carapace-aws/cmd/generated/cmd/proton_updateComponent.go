package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateComponentCmd = &cobra.Command{
	Use:   "update-component",
	Short: "Update a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateComponentCmd).Standalone()

		proton_updateComponentCmd.Flags().String("client-token", "", "The client token for the updated component.")
		proton_updateComponentCmd.Flags().String("deployment-type", "", "The deployment type.")
		proton_updateComponentCmd.Flags().String("description", "", "An optional customer-provided description of the component.")
		proton_updateComponentCmd.Flags().String("name", "", "The name of the component to update.")
		proton_updateComponentCmd.Flags().String("service-instance-name", "", "The name of the service instance that you want to attach this component to.")
		proton_updateComponentCmd.Flags().String("service-name", "", "The name of the service that `serviceInstanceName` is associated with.")
		proton_updateComponentCmd.Flags().String("service-spec", "", "The service spec that you want the component to use to access service inputs.")
		proton_updateComponentCmd.Flags().String("template-file", "", "A path to the Infrastructure as Code (IaC) file describing infrastructure that a custom component provisions.")
		proton_updateComponentCmd.MarkFlagRequired("deployment-type")
		proton_updateComponentCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_updateComponentCmd)
}
