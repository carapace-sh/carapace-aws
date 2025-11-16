package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeHandshakeCmd = &cobra.Command{
	Use:   "describe-handshake",
	Short: "Retrieves information about a previously requested handshake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeHandshakeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_describeHandshakeCmd).Standalone()

		organizations_describeHandshakeCmd.Flags().String("handshake-id", "", "The unique identifier (ID) of the handshake that you want information about.")
		organizations_describeHandshakeCmd.MarkFlagRequired("handshake-id")
	})
	organizationsCmd.AddCommand(organizations_describeHandshakeCmd)
}
