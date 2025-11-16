package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_declineHandshakeCmd = &cobra.Command{
	Use:   "decline-handshake",
	Short: "Declines a handshake request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_declineHandshakeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_declineHandshakeCmd).Standalone()

		organizations_declineHandshakeCmd.Flags().String("handshake-id", "", "The unique identifier (ID) of the handshake that you want to decline.")
		organizations_declineHandshakeCmd.MarkFlagRequired("handshake-id")
	})
	organizationsCmd.AddCommand(organizations_declineHandshakeCmd)
}
