package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createOtaTaskConfigurationCmd = &cobra.Command{
	Use:   "create-ota-task-configuration",
	Short: "Create a configuraiton for the over-the-air (OTA) task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createOtaTaskConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createOtaTaskConfigurationCmd).Standalone()

		iotManagedIntegrations_createOtaTaskConfigurationCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createOtaTaskConfigurationCmd.Flags().String("description", "", "A description of the over-the-air (OTA) task configuration.")
		iotManagedIntegrations_createOtaTaskConfigurationCmd.Flags().String("name", "", "The name of the over-the-air (OTA) task.")
		iotManagedIntegrations_createOtaTaskConfigurationCmd.Flags().String("push-config", "", "Describes the type of configuration used for the over-the-air (OTA) task.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createOtaTaskConfigurationCmd)
}
