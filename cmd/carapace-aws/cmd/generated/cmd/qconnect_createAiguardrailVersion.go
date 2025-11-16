package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAiguardrailVersionCmd = &cobra.Command{
	Use:   "create-aiguardrail-version",
	Short: "Creates an Amazon Q in Connect AI Guardrail version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAiguardrailVersionCmd).Standalone()

	qconnect_createAiguardrailVersionCmd.Flags().String("ai-guardrail-id", "", "The identifier of the Amazon Q in Connect AI Guardrail.")
	qconnect_createAiguardrailVersionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_createAiguardrailVersionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createAiguardrailVersionCmd.Flags().String("modified-time", "", "The time the AI Guardrail was last modified.")
	qconnect_createAiguardrailVersionCmd.MarkFlagRequired("ai-guardrail-id")
	qconnect_createAiguardrailVersionCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_createAiguardrailVersionCmd)
}
