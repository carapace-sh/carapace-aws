package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_cancelHandshakeCmd = &cobra.Command{
	Use:   "cancel-handshake",
	Short: "Cancels a handshake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_cancelHandshakeCmd).Standalone()

	organizations_cancelHandshakeCmd.Flags().String("handshake-id", "", "The unique identifier (ID) of the handshake that you want to cancel.")
	organizations_cancelHandshakeCmd.MarkFlagRequired("handshake-id")
	organizationsCmd.AddCommand(organizations_cancelHandshakeCmd)
}
