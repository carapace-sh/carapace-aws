package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listConnectPeersCmd = &cobra.Command{
	Use:   "list-connect-peers",
	Short: "Returns a list of core network Connect peers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listConnectPeersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_listConnectPeersCmd).Standalone()

		networkmanager_listConnectPeersCmd.Flags().String("connect-attachment-id", "", "The ID of the attachment.")
		networkmanager_listConnectPeersCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_listConnectPeersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_listConnectPeersCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	networkmanagerCmd.AddCommand(networkmanager_listConnectPeersCmd)
}
