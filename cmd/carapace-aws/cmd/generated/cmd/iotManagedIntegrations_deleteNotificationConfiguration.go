package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteNotificationConfigurationCmd = &cobra.Command{
	Use:   "delete-notification-configuration",
	Short: "Deletes a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteNotificationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_deleteNotificationConfigurationCmd).Standalone()

		iotManagedIntegrations_deleteNotificationConfigurationCmd.Flags().String("event-type", "", "The type of event triggering a device notification to the customer-managed destination.")
		iotManagedIntegrations_deleteNotificationConfigurationCmd.MarkFlagRequired("event-type")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteNotificationConfigurationCmd)
}
