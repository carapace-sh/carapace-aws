package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createNotificationConfigurationCmd = &cobra.Command{
	Use:   "create-notification-configuration",
	Short: "Creates a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createNotificationConfigurationCmd).Standalone()

		iotManagedIntegrations_createNotificationConfigurationCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createNotificationConfigurationCmd.Flags().String("destination-name", "", "The name of the destination for the notification configuration.")
		iotManagedIntegrations_createNotificationConfigurationCmd.Flags().String("event-type", "", "The type of event triggering a device notification to the customer-managed destination.")
		iotManagedIntegrations_createNotificationConfigurationCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the notification configuration.")
		iotManagedIntegrations_createNotificationConfigurationCmd.MarkFlagRequired("destination-name")
		iotManagedIntegrations_createNotificationConfigurationCmd.MarkFlagRequired("event-type")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createNotificationConfigurationCmd)
}
