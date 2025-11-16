package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createOtaTaskCmd = &cobra.Command{
	Use:   "create-ota-task",
	Short: "Create an over-the-air (OTA) task to target a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createOtaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createOtaTaskCmd).Standalone()

		iotManagedIntegrations_createOtaTaskCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("description", "", "The description of the over-the-air (OTA) task.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("ota-mechanism", "", "The deployment mechanism for the over-the-air (OTA) task.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("ota-scheduling-config", "", "")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("ota-target-query-string", "", "The query string to add things to the thing group.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("ota-task-execution-retry-config", "", "")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("ota-type", "", "The frequency type for the over-the-air (OTA) task.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("protocol", "", "The connection protocol the over-the-air (OTA) task uses to update the device.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("s3-url", "", "The URL to the Amazon S3 bucket where the over-the-air (OTA) task is stored.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the over-the-air (OTA) task.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("target", "", "The device targeted for the over-the-air (OTA) task.")
		iotManagedIntegrations_createOtaTaskCmd.Flags().String("task-configuration-id", "", "The identifier for the over-the-air (OTA) task configuration.")
		iotManagedIntegrations_createOtaTaskCmd.MarkFlagRequired("ota-type")
		iotManagedIntegrations_createOtaTaskCmd.MarkFlagRequired("s3-url")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createOtaTaskCmd)
}
