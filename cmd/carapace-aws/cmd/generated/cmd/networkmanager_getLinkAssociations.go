package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getLinkAssociationsCmd = &cobra.Command{
	Use:   "get-link-associations",
	Short: "Gets the link associations for a device or a link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getLinkAssociationsCmd).Standalone()

	networkmanager_getLinkAssociationsCmd.Flags().String("device-id", "", "The ID of the device.")
	networkmanager_getLinkAssociationsCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getLinkAssociationsCmd.Flags().String("link-id", "", "The ID of the link.")
	networkmanager_getLinkAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_getLinkAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_getLinkAssociationsCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_getLinkAssociationsCmd)
}
