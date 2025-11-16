package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getNotificationConfigurationCmd = &cobra.Command{
	Use:   "get-notification-configuration",
	Short: "Get a notification configuration for a specified event type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getNotificationConfigurationCmd).Standalone()

	iotManagedIntegrations_getNotificationConfigurationCmd.Flags().String("event-type", "", "The type of event triggering a device notification to the customer-managed destination.")
	iotManagedIntegrations_getNotificationConfigurationCmd.MarkFlagRequired("event-type")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getNotificationConfigurationCmd)
}
