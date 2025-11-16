package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateAiguardrailCmd = &cobra.Command{
	Use:   "update-aiguardrail",
	Short: "Updates an AI Guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateAiguardrailCmd).Standalone()

	qconnect_updateAiguardrailCmd.Flags().String("ai-guardrail-id", "", "The identifier of the Amazon Q in Connect AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_updateAiguardrailCmd.Flags().String("blocked-input-messaging", "", "The message to return when the AI Guardrail blocks a prompt.")
	qconnect_updateAiguardrailCmd.Flags().String("blocked-outputs-messaging", "", "The message to return when the AI Guardrail blocks a model response.")
	qconnect_updateAiguardrailCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_updateAiguardrailCmd.Flags().String("content-policy-config", "", "The content filter policies to configure for the AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("contextual-grounding-policy-config", "", "The contextual grounding policy configuration used to create an AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("description", "", "A description of the AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("sensitive-information-policy-config", "", "The sensitive information policy to configure for the AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("topic-policy-config", "", "The topic policies to configure for the AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("visibility-status", "", "The visibility status of the Amazon Q in Connect AI Guardrail.")
	qconnect_updateAiguardrailCmd.Flags().String("word-policy-config", "", "The word policy you configure for the AI Guardrail.")
	qconnect_updateAiguardrailCmd.MarkFlagRequired("ai-guardrail-id")
	qconnect_updateAiguardrailCmd.MarkFlagRequired("assistant-id")
	qconnect_updateAiguardrailCmd.MarkFlagRequired("blocked-input-messaging")
	qconnect_updateAiguardrailCmd.MarkFlagRequired("blocked-outputs-messaging")
	qconnect_updateAiguardrailCmd.MarkFlagRequired("visibility-status")
	qconnectCmd.AddCommand(qconnect_updateAiguardrailCmd)
}
