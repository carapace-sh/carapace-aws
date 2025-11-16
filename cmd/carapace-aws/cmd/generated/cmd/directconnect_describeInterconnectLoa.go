package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeInterconnectLoaCmd = &cobra.Command{
	Use:   "describe-interconnect-loa",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeInterconnectLoaCmd).Standalone()

	directconnect_describeInterconnectLoaCmd.Flags().String("interconnect-id", "", "The ID of the interconnect.")
	directconnect_describeInterconnectLoaCmd.Flags().String("loa-content-type", "", "The standard media type for the LOA-CFA document.")
	directconnect_describeInterconnectLoaCmd.Flags().String("provider-name", "", "The name of the service provider who establishes connectivity on your behalf.")
	directconnect_describeInterconnectLoaCmd.MarkFlagRequired("interconnect-id")
	directconnectCmd.AddCommand(directconnect_describeInterconnectLoaCmd)
}
