package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_importSshPublicKeyCmd = &cobra.Command{
	Use:   "import-ssh-public-key",
	Short: "Adds a Secure Shell (SSH) public key to a Transfer Family user identified by a `UserName` value assigned to the specific file transfer protocol-enabled server, identified by `ServerId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_importSshPublicKeyCmd).Standalone()

	transfer_importSshPublicKeyCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server.")
	transfer_importSshPublicKeyCmd.Flags().String("ssh-public-key-body", "", "The public key portion of an SSH key pair.")
	transfer_importSshPublicKeyCmd.Flags().String("user-name", "", "The name of the Transfer Family user that is assigned to one or more servers.")
	transfer_importSshPublicKeyCmd.MarkFlagRequired("server-id")
	transfer_importSshPublicKeyCmd.MarkFlagRequired("ssh-public-key-body")
	transfer_importSshPublicKeyCmd.MarkFlagRequired("user-name")
	transferCmd.AddCommand(transfer_importSshPublicKeyCmd)
}
