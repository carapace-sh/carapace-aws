package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_startRemoteDeleteCmd = &cobra.Command{
	Use:   "start-remote-delete",
	Short: "Deletes a file or directory on the remote SFTP server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_startRemoteDeleteCmd).Standalone()

	transfer_startRemoteDeleteCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
	transfer_startRemoteDeleteCmd.Flags().String("delete-path", "", "The absolute path of the file or directory to delete.")
	transfer_startRemoteDeleteCmd.MarkFlagRequired("connector-id")
	transfer_startRemoteDeleteCmd.MarkFlagRequired("delete-path")
	transferCmd.AddCommand(transfer_startRemoteDeleteCmd)
}
