package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_startFileTransferCmd = &cobra.Command{
	Use:   "start-file-transfer",
	Short: "Begins a file transfer between local Amazon Web Services storage and a remote AS2 or SFTP server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_startFileTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_startFileTransferCmd).Standalone()

		transfer_startFileTransferCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
		transfer_startFileTransferCmd.Flags().String("local-directory-path", "", "For an inbound transfer, the `LocaDirectoryPath` specifies the destination for one or more files that are transferred from the partner's SFTP server.")
		transfer_startFileTransferCmd.Flags().String("remote-directory-path", "", "For an outbound transfer, the `RemoteDirectoryPath` specifies the destination for one or more files that are transferred to the partner's SFTP server.")
		transfer_startFileTransferCmd.Flags().String("retrieve-file-paths", "", "One or more source paths for the partner's SFTP server.")
		transfer_startFileTransferCmd.Flags().String("send-file-paths", "", "One or more source paths for the Amazon S3 storage.")
		transfer_startFileTransferCmd.MarkFlagRequired("connector-id")
	})
	transferCmd.AddCommand(transfer_startFileTransferCmd)
}
