package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_stopServerCmd = &cobra.Command{
	Use:   "stop-server",
	Short: "Changes the state of a file transfer protocol-enabled server from `ONLINE` to `OFFLINE`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_stopServerCmd).Standalone()

	transfer_stopServerCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that you stopped.")
	transfer_stopServerCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_stopServerCmd)
}
