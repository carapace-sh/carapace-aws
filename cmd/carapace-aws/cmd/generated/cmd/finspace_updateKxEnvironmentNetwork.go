package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxEnvironmentNetworkCmd = &cobra.Command{
	Use:   "update-kx-environment-network",
	Short: "Updates environment network to connect to your internal network by using a transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxEnvironmentNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_updateKxEnvironmentNetworkCmd).Standalone()

		finspace_updateKxEnvironmentNetworkCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspace_updateKxEnvironmentNetworkCmd.Flags().String("custom-dnsconfiguration", "", "A list of DNS server name and server IP.")
		finspace_updateKxEnvironmentNetworkCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_updateKxEnvironmentNetworkCmd.Flags().String("transit-gateway-configuration", "", "Specifies the transit gateway and network configuration to connect the kdb environment to an internal network.")
		finspace_updateKxEnvironmentNetworkCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_updateKxEnvironmentNetworkCmd)
}
