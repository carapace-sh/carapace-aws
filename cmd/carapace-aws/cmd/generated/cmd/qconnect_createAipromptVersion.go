package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAipromptVersionCmd = &cobra.Command{
	Use:   "create-aiprompt-version",
	Short: "Creates an Amazon Q in Connect AI Prompt version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAipromptVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createAipromptVersionCmd).Standalone()

		qconnect_createAipromptVersionCmd.Flags().String("ai-prompt-id", "", "The identifier of the Amazon Q in Connect AI prompt.")
		qconnect_createAipromptVersionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_createAipromptVersionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createAipromptVersionCmd.Flags().String("modified-time", "", "The time the AI Prompt was last modified.")
		qconnect_createAipromptVersionCmd.MarkFlagRequired("ai-prompt-id")
		qconnect_createAipromptVersionCmd.MarkFlagRequired("assistant-id")
	})
	qconnectCmd.AddCommand(qconnect_createAipromptVersionCmd)
}
