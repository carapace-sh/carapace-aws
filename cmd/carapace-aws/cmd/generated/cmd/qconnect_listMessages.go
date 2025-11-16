package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listMessagesCmd = &cobra.Command{
	Use:   "list-messages",
	Short: "Lists messages on an Amazon Q in Connect session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listMessagesCmd).Standalone()

	qconnect_listMessagesCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listMessagesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listMessagesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listMessagesCmd.Flags().String("session-id", "", "The identifier of the Amazon Q in Connect session.")
	qconnect_listMessagesCmd.MarkFlagRequired("assistant-id")
	qconnect_listMessagesCmd.MarkFlagRequired("session-id")
	qconnectCmd.AddCommand(qconnect_listMessagesCmd)
}
