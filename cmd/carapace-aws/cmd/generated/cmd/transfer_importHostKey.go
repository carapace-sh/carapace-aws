package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_importHostKeyCmd = &cobra.Command{
	Use:   "import-host-key",
	Short: "Adds a host key to the server that's specified by the `ServerId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_importHostKeyCmd).Standalone()

	transfer_importHostKeyCmd.Flags().String("description", "", "The text description that identifies this host key.")
	transfer_importHostKeyCmd.Flags().String("host-key-body", "", "The private key portion of an SSH key pair.")
	transfer_importHostKeyCmd.Flags().String("server-id", "", "The identifier of the server that contains the host key that you are importing.")
	transfer_importHostKeyCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for host keys.")
	transfer_importHostKeyCmd.MarkFlagRequired("host-key-body")
	transfer_importHostKeyCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_importHostKeyCmd)
}
