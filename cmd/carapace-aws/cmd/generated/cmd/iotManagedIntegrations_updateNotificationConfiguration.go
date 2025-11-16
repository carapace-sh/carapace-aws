package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateNotificationConfigurationCmd = &cobra.Command{
	Use:   "update-notification-configuration",
	Short: "Update a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateNotificationConfigurationCmd).Standalone()

	iotManagedIntegrations_updateNotificationConfigurationCmd.Flags().String("destination-name", "", "The name of the destination for the notification configuration.")
	iotManagedIntegrations_updateNotificationConfigurationCmd.Flags().String("event-type", "", "The type of event triggering a device notification to the customer-managed destination.")
	iotManagedIntegrations_updateNotificationConfigurationCmd.MarkFlagRequired("destination-name")
	iotManagedIntegrations_updateNotificationConfigurationCmd.MarkFlagRequired("event-type")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateNotificationConfigurationCmd)
}
