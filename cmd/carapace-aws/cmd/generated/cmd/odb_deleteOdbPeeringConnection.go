package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_deleteOdbPeeringConnectionCmd = &cobra.Command{
	Use:   "delete-odb-peering-connection",
	Short: "Deletes an ODB peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_deleteOdbPeeringConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_deleteOdbPeeringConnectionCmd).Standalone()

		odb_deleteOdbPeeringConnectionCmd.Flags().String("odb-peering-connection-id", "", "The unique identifier of the ODB peering connection to delete.")
		odb_deleteOdbPeeringConnectionCmd.MarkFlagRequired("odb-peering-connection-id")
	})
	odbCmd.AddCommand(odb_deleteOdbPeeringConnectionCmd)
}
