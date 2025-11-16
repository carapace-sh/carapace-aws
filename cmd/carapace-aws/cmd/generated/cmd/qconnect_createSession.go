package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createSessionCmd = &cobra.Command{
	Use:   "create-session",
	Short: "Creates a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createSessionCmd).Standalone()

		qconnect_createSessionCmd.Flags().String("ai-agent-configuration", "", "The configuration of the AI Agents (mapped by AI Agent Type to AI Agent version) that should be used by Amazon Q in Connect for this Session.")
		qconnect_createSessionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_createSessionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createSessionCmd.Flags().String("contact-arn", "", "The Amazon Resource Name (ARN) of the email contact in Amazon Connect.")
		qconnect_createSessionCmd.Flags().String("description", "", "The description.")
		qconnect_createSessionCmd.Flags().String("name", "", "The name of the session.")
		qconnect_createSessionCmd.Flags().String("tag-filter", "", "An object that can be used to specify Tag conditions.")
		qconnect_createSessionCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_createSessionCmd.MarkFlagRequired("assistant-id")
		qconnect_createSessionCmd.MarkFlagRequired("name")
	})
	qconnectCmd.AddCommand(qconnect_createSessionCmd)
}
