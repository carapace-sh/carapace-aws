package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_registerCustomEndpointCmd = &cobra.Command{
	Use:   "register-custom-endpoint",
	Short: "Customers can request IoT managed integrations to manage the server trust for them or bring their own external server trusts for the custom domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_registerCustomEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_registerCustomEndpointCmd).Standalone()

	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_registerCustomEndpointCmd)
}
