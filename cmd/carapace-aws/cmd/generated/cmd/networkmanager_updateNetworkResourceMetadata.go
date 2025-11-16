package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateNetworkResourceMetadataCmd = &cobra.Command{
	Use:   "update-network-resource-metadata",
	Short: "Updates the resource metadata for the specified global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateNetworkResourceMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_updateNetworkResourceMetadataCmd).Standalone()

		networkmanager_updateNetworkResourceMetadataCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_updateNetworkResourceMetadataCmd.Flags().String("metadata", "", "The resource metadata.")
		networkmanager_updateNetworkResourceMetadataCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		networkmanager_updateNetworkResourceMetadataCmd.MarkFlagRequired("global-network-id")
		networkmanager_updateNetworkResourceMetadataCmd.MarkFlagRequired("metadata")
		networkmanager_updateNetworkResourceMetadataCmd.MarkFlagRequired("resource-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_updateNetworkResourceMetadataCmd)
}
