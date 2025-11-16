package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_createOdbPeeringConnectionCmd = &cobra.Command{
	Use:   "create-odb-peering-connection",
	Short: "Creates a peering connection between an ODB network and a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_createOdbPeeringConnectionCmd).Standalone()

	odb_createOdbPeeringConnectionCmd.Flags().String("client-token", "", "The client token for the ODB peering connection request.")
	odb_createOdbPeeringConnectionCmd.Flags().String("display-name", "", "The display name for the ODB peering connection.")
	odb_createOdbPeeringConnectionCmd.Flags().String("odb-network-id", "", "The unique identifier of the ODB network that initiates the peering connection.")
	odb_createOdbPeeringConnectionCmd.Flags().String("peer-network-cidrs-to-be-added", "", "A list of CIDR blocks to add to the peering connection.")
	odb_createOdbPeeringConnectionCmd.Flags().String("peer-network-id", "", "The unique identifier of the peer network.")
	odb_createOdbPeeringConnectionCmd.Flags().String("tags", "", "The tags to assign to the ODB peering connection.")
	odb_createOdbPeeringConnectionCmd.MarkFlagRequired("odb-network-id")
	odb_createOdbPeeringConnectionCmd.MarkFlagRequired("peer-network-id")
	odbCmd.AddCommand(odb_createOdbPeeringConnectionCmd)
}
