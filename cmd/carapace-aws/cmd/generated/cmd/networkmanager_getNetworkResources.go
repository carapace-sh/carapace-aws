package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getNetworkResourcesCmd = &cobra.Command{
	Use:   "get-network-resources",
	Short: "Describes the network resources for the specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getNetworkResourcesCmd).Standalone()

	networkmanager_getNetworkResourcesCmd.Flags().String("account-id", "", "The Amazon Web Services account ID.")
	networkmanager_getNetworkResourcesCmd.Flags().String("aws-region", "", "The Amazon Web Services Region.")
	networkmanager_getNetworkResourcesCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_getNetworkResourcesCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getNetworkResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_getNetworkResourcesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_getNetworkResourcesCmd.Flags().String("registered-gateway-arn", "", "The ARN of the gateway.")
	networkmanager_getNetworkResourcesCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	networkmanager_getNetworkResourcesCmd.Flags().String("resource-type", "", "The resource type.")
	networkmanager_getNetworkResourcesCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_getNetworkResourcesCmd)
}
