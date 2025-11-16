package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteHostKeyCmd = &cobra.Command{
	Use:   "delete-host-key",
	Short: "Deletes the host key that's specified in the `HostKeyId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteHostKeyCmd).Standalone()

	transfer_deleteHostKeyCmd.Flags().String("host-key-id", "", "The identifier of the host key that you are deleting.")
	transfer_deleteHostKeyCmd.Flags().String("server-id", "", "The identifier of the server that contains the host key that you are deleting.")
	transfer_deleteHostKeyCmd.MarkFlagRequired("host-key-id")
	transfer_deleteHostKeyCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_deleteHostKeyCmd)
}
