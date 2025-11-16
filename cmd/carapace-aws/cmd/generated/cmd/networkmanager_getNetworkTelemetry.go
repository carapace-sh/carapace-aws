package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getNetworkTelemetryCmd = &cobra.Command{
	Use:   "get-network-telemetry",
	Short: "Gets the network telemetry of the specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getNetworkTelemetryCmd).Standalone()

	networkmanager_getNetworkTelemetryCmd.Flags().String("account-id", "", "The Amazon Web Services account ID.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("aws-region", "", "The Amazon Web Services Region.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("registered-gateway-arn", "", "The ARN of the gateway.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	networkmanager_getNetworkTelemetryCmd.Flags().String("resource-type", "", "The resource type.")
	networkmanager_getNetworkTelemetryCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_getNetworkTelemetryCmd)
}
