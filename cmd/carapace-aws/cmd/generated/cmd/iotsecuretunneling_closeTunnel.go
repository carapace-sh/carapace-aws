package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_closeTunnelCmd = &cobra.Command{
	Use:   "close-tunnel",
	Short: "Closes a tunnel identified by the unique tunnel id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_closeTunnelCmd).Standalone()

	iotsecuretunneling_closeTunnelCmd.Flags().String("delete", "", "When set to true, IoT Secure Tunneling deletes the tunnel data immediately.")
	iotsecuretunneling_closeTunnelCmd.Flags().String("tunnel-id", "", "The ID of the tunnel to close.")
	iotsecuretunneling_closeTunnelCmd.MarkFlagRequired("tunnel-id")
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_closeTunnelCmd)
}
