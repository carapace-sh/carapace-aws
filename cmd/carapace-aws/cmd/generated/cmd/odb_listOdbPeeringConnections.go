package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listOdbPeeringConnectionsCmd = &cobra.Command{
	Use:   "list-odb-peering-connections",
	Short: "Lists all ODB peering connections or those associated with a specific ODB network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listOdbPeeringConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listOdbPeeringConnectionsCmd).Standalone()

		odb_listOdbPeeringConnectionsCmd.Flags().String("max-results", "", "The maximum number of ODB peering connections to return in the response.")
		odb_listOdbPeeringConnectionsCmd.Flags().String("next-token", "", "The pagination token for the next page of ODB peering connections.")
		odb_listOdbPeeringConnectionsCmd.Flags().String("odb-network-id", "", "The identifier of the ODB network to list peering connections for.")
	})
	odbCmd.AddCommand(odb_listOdbPeeringConnectionsCmd)
}
