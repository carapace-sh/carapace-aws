package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getNextMessageCmd = &cobra.Command{
	Use:   "get-next-message",
	Short: "Retrieves next message on an Amazon Q in Connect session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getNextMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getNextMessageCmd).Standalone()

		qconnect_getNextMessageCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_getNextMessageCmd.Flags().String("next-message-token", "", "The token for the next message.")
		qconnect_getNextMessageCmd.Flags().String("session-id", "", "The identifier of the Amazon Q in Connect session.")
		qconnect_getNextMessageCmd.MarkFlagRequired("assistant-id")
		qconnect_getNextMessageCmd.MarkFlagRequired("next-message-token")
		qconnect_getNextMessageCmd.MarkFlagRequired("session-id")
	})
	qconnectCmd.AddCommand(qconnect_getNextMessageCmd)
}
