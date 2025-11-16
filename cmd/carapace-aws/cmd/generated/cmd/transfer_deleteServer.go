package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteServerCmd = &cobra.Command{
	Use:   "delete-server",
	Short: "Deletes the file transfer protocol-enabled server that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteServerCmd).Standalone()

	transfer_deleteServerCmd.Flags().String("server-id", "", "A unique system-assigned identifier for a server instance.")
	transfer_deleteServerCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_deleteServerCmd)
}
