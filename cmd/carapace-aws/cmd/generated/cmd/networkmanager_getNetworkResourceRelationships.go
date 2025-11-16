package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getNetworkResourceRelationshipsCmd = &cobra.Command{
	Use:   "get-network-resource-relationships",
	Short: "Gets the network resource relationships for the specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getNetworkResourceRelationshipsCmd).Standalone()

	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("aws-region", "", "The Amazon Web Services Region.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("core-network-id", "", "The ID of a core network.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("registered-gateway-arn", "", "The ARN of the registered gateway.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("resource-arn", "", "The ARN of the gateway.")
	networkmanager_getNetworkResourceRelationshipsCmd.Flags().String("resource-type", "", "The resource type.")
	networkmanager_getNetworkResourceRelationshipsCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_getNetworkResourceRelationshipsCmd)
}
