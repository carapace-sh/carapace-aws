package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateSessionCmd = &cobra.Command{
	Use:   "update-session",
	Short: "Updates a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateSessionCmd).Standalone()

	qconnect_updateSessionCmd.Flags().String("ai-agent-configuration", "", "The configuration of the AI Agents (mapped by AI Agent Type to AI Agent version) that should be used by Amazon Q in Connect for this Session.")
	qconnect_updateSessionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_updateSessionCmd.Flags().String("description", "", "The description.")
	qconnect_updateSessionCmd.Flags().String("session-id", "", "The identifier of the session.")
	qconnect_updateSessionCmd.Flags().String("tag-filter", "", "An object that can be used to specify Tag conditions.")
	qconnect_updateSessionCmd.MarkFlagRequired("assistant-id")
	qconnect_updateSessionCmd.MarkFlagRequired("session-id")
	qconnectCmd.AddCommand(qconnect_updateSessionCmd)
}
