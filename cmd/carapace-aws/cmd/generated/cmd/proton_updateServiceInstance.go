package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServiceInstanceCmd = &cobra.Command{
	Use:   "update-service-instance",
	Short: "Update a service instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServiceInstanceCmd).Standalone()

	proton_updateServiceInstanceCmd.Flags().String("client-token", "", "The client token of the service instance to update.")
	proton_updateServiceInstanceCmd.Flags().String("deployment-type", "", "The deployment type.")
	proton_updateServiceInstanceCmd.Flags().String("name", "", "The name of the service instance to update.")
	proton_updateServiceInstanceCmd.Flags().String("service-name", "", "The name of the service that the service instance belongs to.")
	proton_updateServiceInstanceCmd.Flags().String("spec", "", "The formatted specification that defines the service instance update.")
	proton_updateServiceInstanceCmd.Flags().String("template-major-version", "", "The major version of the service template to update.")
	proton_updateServiceInstanceCmd.Flags().String("template-minor-version", "", "The minor version of the service template to update.")
	proton_updateServiceInstanceCmd.MarkFlagRequired("deployment-type")
	proton_updateServiceInstanceCmd.MarkFlagRequired("name")
	proton_updateServiceInstanceCmd.MarkFlagRequired("service-name")
	protonCmd.AddCommand(proton_updateServiceInstanceCmd)
}
