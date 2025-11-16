package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_updateOdbPeeringConnectionCmd = &cobra.Command{
	Use:   "update-odb-peering-connection",
	Short: "Modifies the settings of an Oracle Database@Amazon Web Services peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_updateOdbPeeringConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_updateOdbPeeringConnectionCmd).Standalone()

		odb_updateOdbPeeringConnectionCmd.Flags().String("display-name", "", "A new display name for the peering connection.")
		odb_updateOdbPeeringConnectionCmd.Flags().String("odb-peering-connection-id", "", "The identifier of the Oracle Database@Amazon Web Services peering connection to update.")
		odb_updateOdbPeeringConnectionCmd.Flags().String("peer-network-cidrs-to-be-added", "", "A list of CIDR blocks to add to the peering connection.")
		odb_updateOdbPeeringConnectionCmd.Flags().String("peer-network-cidrs-to-be-removed", "", "A list of CIDR blocks to remove from the peering connection.")
		odb_updateOdbPeeringConnectionCmd.MarkFlagRequired("odb-peering-connection-id")
	})
	odbCmd.AddCommand(odb_updateOdbPeeringConnectionCmd)
}
