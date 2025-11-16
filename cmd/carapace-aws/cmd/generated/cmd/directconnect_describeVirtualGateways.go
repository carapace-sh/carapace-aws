package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeVirtualGatewaysCmd = &cobra.Command{
	Use:   "describe-virtual-gateways",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeVirtualGatewaysCmd).Standalone()

	directconnectCmd.AddCommand(directconnect_describeVirtualGatewaysCmd)
}
