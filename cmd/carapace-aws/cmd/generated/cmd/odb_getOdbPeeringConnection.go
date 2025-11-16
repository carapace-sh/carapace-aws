package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getOdbPeeringConnectionCmd = &cobra.Command{
	Use:   "get-odb-peering-connection",
	Short: "Retrieves information about an ODB peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getOdbPeeringConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_getOdbPeeringConnectionCmd).Standalone()

		odb_getOdbPeeringConnectionCmd.Flags().String("odb-peering-connection-id", "", "The unique identifier of the ODB peering connection to retrieve information about.")
		odb_getOdbPeeringConnectionCmd.MarkFlagRequired("odb-peering-connection-id")
	})
	odbCmd.AddCommand(odb_getOdbPeeringConnectionCmd)
}
