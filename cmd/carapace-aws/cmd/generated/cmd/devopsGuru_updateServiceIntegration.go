package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_updateServiceIntegrationCmd = &cobra.Command{
	Use:   "update-service-integration",
	Short: "Enables or disables integration with a service that can be integrated with DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_updateServiceIntegrationCmd).Standalone()

	devopsGuru_updateServiceIntegrationCmd.Flags().String("service-integration", "", "An `IntegratedServiceConfig` object used to specify the integrated service you want to update, and whether you want to update it to enabled or disabled.")
	devopsGuru_updateServiceIntegrationCmd.MarkFlagRequired("service-integration")
	devopsGuruCmd.AddCommand(devopsGuru_updateServiceIntegrationCmd)
}
