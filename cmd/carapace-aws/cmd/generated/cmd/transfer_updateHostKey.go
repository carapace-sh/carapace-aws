package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateHostKeyCmd = &cobra.Command{
	Use:   "update-host-key",
	Short: "Updates the description for the host key that's specified by the `ServerId` and `HostKeyId` parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateHostKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_updateHostKeyCmd).Standalone()

		transfer_updateHostKeyCmd.Flags().String("description", "", "An updated description for the host key.")
		transfer_updateHostKeyCmd.Flags().String("host-key-id", "", "The identifier of the host key that you are updating.")
		transfer_updateHostKeyCmd.Flags().String("server-id", "", "The identifier of the server that contains the host key that you are updating.")
		transfer_updateHostKeyCmd.MarkFlagRequired("description")
		transfer_updateHostKeyCmd.MarkFlagRequired("host-key-id")
		transfer_updateHostKeyCmd.MarkFlagRequired("server-id")
	})
	transferCmd.AddCommand(transfer_updateHostKeyCmd)
}
