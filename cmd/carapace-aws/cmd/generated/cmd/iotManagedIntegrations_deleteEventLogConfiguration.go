package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteEventLogConfigurationCmd = &cobra.Command{
	Use:   "delete-event-log-configuration",
	Short: "Delete an event log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteEventLogConfigurationCmd).Standalone()

	iotManagedIntegrations_deleteEventLogConfigurationCmd.Flags().String("id", "", "The identifier of the event log configuration.")
	iotManagedIntegrations_deleteEventLogConfigurationCmd.MarkFlagRequired("id")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteEventLogConfigurationCmd)
}
