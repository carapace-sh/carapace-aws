package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_openTunnelCmd = &cobra.Command{
	Use:   "open-tunnel",
	Short: "Creates a new tunnel, and returns two client access tokens for clients to use to connect to the IoT Secure Tunneling proxy server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_openTunnelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunneling_openTunnelCmd).Standalone()

		iotsecuretunneling_openTunnelCmd.Flags().String("description", "", "A short text description of the tunnel.")
		iotsecuretunneling_openTunnelCmd.Flags().String("destination-config", "", "The destination configuration for the OpenTunnel request.")
		iotsecuretunneling_openTunnelCmd.Flags().String("tags", "", "A collection of tag metadata.")
		iotsecuretunneling_openTunnelCmd.Flags().String("timeout-config", "", "Timeout configuration for a tunnel.")
	})
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_openTunnelCmd)
}
