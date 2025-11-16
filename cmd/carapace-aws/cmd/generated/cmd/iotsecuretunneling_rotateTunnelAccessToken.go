package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsecuretunneling_rotateTunnelAccessTokenCmd = &cobra.Command{
	Use:   "rotate-tunnel-access-token",
	Short: "Revokes the current client access token (CAT) and returns new CAT for clients to use when reconnecting to secure tunneling to access the same tunnel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsecuretunneling_rotateTunnelAccessTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsecuretunneling_rotateTunnelAccessTokenCmd).Standalone()

		iotsecuretunneling_rotateTunnelAccessTokenCmd.Flags().String("client-mode", "", "The mode of the client that will use the client token, which can be either the source or destination, or both source and destination.")
		iotsecuretunneling_rotateTunnelAccessTokenCmd.Flags().String("destination-config", "", "")
		iotsecuretunneling_rotateTunnelAccessTokenCmd.Flags().String("tunnel-id", "", "The tunnel for which you want to rotate the access tokens.")
		iotsecuretunneling_rotateTunnelAccessTokenCmd.MarkFlagRequired("client-mode")
		iotsecuretunneling_rotateTunnelAccessTokenCmd.MarkFlagRequired("tunnel-id")
	})
	iotsecuretunnelingCmd.AddCommand(iotsecuretunneling_rotateTunnelAccessTokenCmd)
}
