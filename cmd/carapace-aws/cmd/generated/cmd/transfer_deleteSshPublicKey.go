package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteSshPublicKeyCmd = &cobra.Command{
	Use:   "delete-ssh-public-key",
	Short: "Deletes a user's Secure Shell (SSH) public key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteSshPublicKeyCmd).Standalone()

	transfer_deleteSshPublicKeyCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a file transfer protocol-enabled server instance that has the user assigned to it.")
	transfer_deleteSshPublicKeyCmd.Flags().String("ssh-public-key-id", "", "A unique identifier used to reference your user's specific SSH key.")
	transfer_deleteSshPublicKeyCmd.Flags().String("user-name", "", "A unique string that identifies a user whose public key is being deleted.")
	transfer_deleteSshPublicKeyCmd.MarkFlagRequired("server-id")
	transfer_deleteSshPublicKeyCmd.MarkFlagRequired("ssh-public-key-id")
	transfer_deleteSshPublicKeyCmd.MarkFlagRequired("user-name")
	transferCmd.AddCommand(transfer_deleteSshPublicKeyCmd)
}
