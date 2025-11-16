package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Changes the state of a file transfer protocol-enabled server from `OFFLINE` to `ONLINE`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_startServerCmd).Standalone()

	transfer_startServerCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that you start.")
	transfer_startServerCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_startServerCmd)
}
