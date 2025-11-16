package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getHubConfigurationCmd = &cobra.Command{
	Use:   "get-hub-configuration",
	Short: "Get a hub configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getHubConfigurationCmd).Standalone()

	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getHubConfigurationCmd)
}
