package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_startRemoteMoveCmd = &cobra.Command{
	Use:   "start-remote-move",
	Short: "Moves or renames a file or directory on the remote SFTP server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_startRemoteMoveCmd).Standalone()

	transfer_startRemoteMoveCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
	transfer_startRemoteMoveCmd.Flags().String("source-path", "", "The absolute path of the file or directory to move or rename.")
	transfer_startRemoteMoveCmd.Flags().String("target-path", "", "The absolute path for the target of the move/rename operation.")
	transfer_startRemoteMoveCmd.MarkFlagRequired("connector-id")
	transfer_startRemoteMoveCmd.MarkFlagRequired("source-path")
	transfer_startRemoteMoveCmd.MarkFlagRequired("target-path")
	transferCmd.AddCommand(transfer_startRemoteMoveCmd)
}
