package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_acceptHandshakeCmd = &cobra.Command{
	Use:   "accept-handshake",
	Short: "Sends a response to the originator of a handshake agreeing to the action proposed by the handshake request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_acceptHandshakeCmd).Standalone()

	organizations_acceptHandshakeCmd.Flags().String("handshake-id", "", "The unique identifier (ID) of the handshake that you want to accept.")
	organizations_acceptHandshakeCmd.MarkFlagRequired("handshake-id")
	organizationsCmd.AddCommand(organizations_acceptHandshakeCmd)
}
