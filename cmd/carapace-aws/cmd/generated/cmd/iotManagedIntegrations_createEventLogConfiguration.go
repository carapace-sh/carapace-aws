package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createEventLogConfigurationCmd = &cobra.Command{
	Use:   "create-event-log-configuration",
	Short: "Set the event log configuration for the account, resource type, or specific resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createEventLogConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createEventLogConfigurationCmd).Standalone()

		iotManagedIntegrations_createEventLogConfigurationCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createEventLogConfigurationCmd.Flags().String("event-log-level", "", "The logging level for the event log configuration.")
		iotManagedIntegrations_createEventLogConfigurationCmd.Flags().String("resource-id", "", "The identifier of the resource for the event log configuration.")
		iotManagedIntegrations_createEventLogConfigurationCmd.Flags().String("resource-type", "", "The type of resource for the event log configuration.")
		iotManagedIntegrations_createEventLogConfigurationCmd.MarkFlagRequired("event-log-level")
		iotManagedIntegrations_createEventLogConfigurationCmd.MarkFlagRequired("resource-type")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createEventLogConfigurationCmd)
}
