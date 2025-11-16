package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getEventLogConfigurationCmd = &cobra.Command{
	Use:   "get-event-log-configuration",
	Short: "Get an event log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getEventLogConfigurationCmd).Standalone()

	iotManagedIntegrations_getEventLogConfigurationCmd.Flags().String("id", "", "The identifier of the event log configuration.")
	iotManagedIntegrations_getEventLogConfigurationCmd.MarkFlagRequired("id")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getEventLogConfigurationCmd)
}
