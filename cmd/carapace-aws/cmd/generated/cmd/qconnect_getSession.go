package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Retrieves information for a specified session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getSessionCmd).Standalone()

	qconnect_getSessionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_getSessionCmd.Flags().String("session-id", "", "The identifier of the session.")
	qconnect_getSessionCmd.MarkFlagRequired("assistant-id")
	qconnect_getSessionCmd.MarkFlagRequired("session-id")
	qconnectCmd.AddCommand(qconnect_getSessionCmd)
}
