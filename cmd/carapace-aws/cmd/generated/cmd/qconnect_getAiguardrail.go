package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getAiguardrailCmd = &cobra.Command{
	Use:   "get-aiguardrail",
	Short: "Gets the Amazon Q in Connect AI Guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getAiguardrailCmd).Standalone()

	qconnect_getAiguardrailCmd.Flags().String("ai-guardrail-id", "", "The identifier of the Amazon Q in Connect AI Guardrail.")
	qconnect_getAiguardrailCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_getAiguardrailCmd.MarkFlagRequired("ai-guardrail-id")
	qconnect_getAiguardrailCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_getAiguardrailCmd)
}
