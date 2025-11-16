package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getOdbNetworkCmd = &cobra.Command{
	Use:   "get-odb-network",
	Short: "Returns information about the specified ODB network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getOdbNetworkCmd).Standalone()

	odb_getOdbNetworkCmd.Flags().String("odb-network-id", "", "The unique identifier of the ODB network.")
	odb_getOdbNetworkCmd.MarkFlagRequired("odb-network-id")
	odbCmd.AddCommand(odb_getOdbNetworkCmd)
}
