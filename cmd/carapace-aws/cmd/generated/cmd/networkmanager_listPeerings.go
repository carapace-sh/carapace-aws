package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listPeeringsCmd = &cobra.Command{
	Use:   "list-peerings",
	Short: "Lists the peerings for a core network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listPeeringsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_listPeeringsCmd).Standalone()

		networkmanager_listPeeringsCmd.Flags().String("core-network-id", "", "The ID of a core network.")
		networkmanager_listPeeringsCmd.Flags().String("edge-location", "", "Returns a list edge locations for the")
		networkmanager_listPeeringsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		networkmanager_listPeeringsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		networkmanager_listPeeringsCmd.Flags().String("peering-type", "", "Returns a list of a peering requests.")
		networkmanager_listPeeringsCmd.Flags().String("state", "", "Returns a list of the peering request states.")
	})
	networkmanagerCmd.AddCommand(networkmanager_listPeeringsCmd)
}
