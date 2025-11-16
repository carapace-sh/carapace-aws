package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeLoaCmd = &cobra.Command{
	Use:   "describe-loa",
	Short: "Gets the LOA-CFA for a connection, interconnect, or link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeLoaCmd).Standalone()

	directconnect_describeLoaCmd.Flags().String("connection-id", "", "The ID of a connection, LAG, or interconnect.")
	directconnect_describeLoaCmd.Flags().String("loa-content-type", "", "The standard media type for the LOA-CFA document.")
	directconnect_describeLoaCmd.Flags().String("provider-name", "", "The name of the service provider who establishes connectivity on your behalf.")
	directconnect_describeLoaCmd.MarkFlagRequired("connection-id")
	directconnectCmd.AddCommand(directconnect_describeLoaCmd)
}
