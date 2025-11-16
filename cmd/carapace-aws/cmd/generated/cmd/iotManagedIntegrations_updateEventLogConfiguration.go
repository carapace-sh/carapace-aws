package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateEventLogConfigurationCmd = &cobra.Command{
	Use:   "update-event-log-configuration",
	Short: "Update an event log configuration by log configuration ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateEventLogConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_updateEventLogConfigurationCmd).Standalone()

		iotManagedIntegrations_updateEventLogConfigurationCmd.Flags().String("event-log-level", "", "The log level for the event in terms of severity.")
		iotManagedIntegrations_updateEventLogConfigurationCmd.Flags().String("id", "", "The log configuration id.")
		iotManagedIntegrations_updateEventLogConfigurationCmd.MarkFlagRequired("event-log-level")
		iotManagedIntegrations_updateEventLogConfigurationCmd.MarkFlagRequired("id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateEventLogConfigurationCmd)
}
