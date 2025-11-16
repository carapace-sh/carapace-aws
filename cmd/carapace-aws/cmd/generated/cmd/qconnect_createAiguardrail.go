package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAiguardrailCmd = &cobra.Command{
	Use:   "create-aiguardrail",
	Short: "Creates an Amazon Q in Connect AI Guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAiguardrailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createAiguardrailCmd).Standalone()

		qconnect_createAiguardrailCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_createAiguardrailCmd.Flags().String("blocked-input-messaging", "", "The message to return when the AI Guardrail blocks a prompt.")
		qconnect_createAiguardrailCmd.Flags().String("blocked-outputs-messaging", "", "The message to return when the AI Guardrail blocks a model response.")
		qconnect_createAiguardrailCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createAiguardrailCmd.Flags().String("content-policy-config", "", "The content filter policies to configure for the AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("contextual-grounding-policy-config", "", "The contextual grounding policy configuration used to create an AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("description", "", "A description of the AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("name", "", "The name of the AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("sensitive-information-policy-config", "", "The sensitive information policy to configure for the AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_createAiguardrailCmd.Flags().String("topic-policy-config", "", "The topic policies to configure for the AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("visibility-status", "", "The visibility status of the AI Guardrail.")
		qconnect_createAiguardrailCmd.Flags().String("word-policy-config", "", "The word policy you configure for the AI Guardrail.")
		qconnect_createAiguardrailCmd.MarkFlagRequired("assistant-id")
		qconnect_createAiguardrailCmd.MarkFlagRequired("blocked-input-messaging")
		qconnect_createAiguardrailCmd.MarkFlagRequired("blocked-outputs-messaging")
		qconnect_createAiguardrailCmd.MarkFlagRequired("name")
		qconnect_createAiguardrailCmd.MarkFlagRequired("visibility-status")
	})
	qconnectCmd.AddCommand(qconnect_createAiguardrailCmd)
}
