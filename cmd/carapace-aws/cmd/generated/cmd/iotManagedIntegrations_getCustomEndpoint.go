package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getCustomEndpointCmd = &cobra.Command{
	Use:   "get-custom-endpoint",
	Short: "Returns the IoT managed integrations custom endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getCustomEndpointCmd).Standalone()

	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getCustomEndpointCmd)
}
