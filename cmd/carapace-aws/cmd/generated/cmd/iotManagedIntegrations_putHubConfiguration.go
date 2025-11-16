package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_putHubConfigurationCmd = &cobra.Command{
	Use:   "put-hub-configuration",
	Short: "Update a hub configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_putHubConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_putHubConfigurationCmd).Standalone()

		iotManagedIntegrations_putHubConfigurationCmd.Flags().String("hub-token-timer-expiry-setting-in-seconds", "", "A user-defined integer value that represents the hub token timer expiry setting in seconds.")
		iotManagedIntegrations_putHubConfigurationCmd.MarkFlagRequired("hub-token-timer-expiry-setting-in-seconds")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_putHubConfigurationCmd)
}
