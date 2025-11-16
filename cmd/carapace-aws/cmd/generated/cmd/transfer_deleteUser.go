package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes the user belonging to a file transfer protocol-enabled server you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteUserCmd).Standalone()

		transfer_deleteUserCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance that has the user assigned to it.")
		transfer_deleteUserCmd.Flags().String("user-name", "", "A unique string that identifies a user that is being deleted from a server.")
		transfer_deleteUserCmd.MarkFlagRequired("server-id")
		transfer_deleteUserCmd.MarkFlagRequired("user-name")
	})
	transferCmd.AddCommand(transfer_deleteUserCmd)
}
