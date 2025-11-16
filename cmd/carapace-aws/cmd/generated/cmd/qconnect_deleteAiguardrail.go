package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAiguardrailCmd = &cobra.Command{
	Use:   "delete-aiguardrail",
	Short: "Deletes an Amazon Q in Connect AI Guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAiguardrailCmd).Standalone()

	qconnect_deleteAiguardrailCmd.Flags().String("ai-guardrail-id", "", "The identifier of the Amazon Q in Connect AI Guardrail.")
	qconnect_deleteAiguardrailCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_deleteAiguardrailCmd.MarkFlagRequired("ai-guardrail-id")
	qconnect_deleteAiguardrailCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_deleteAiguardrailCmd)
}
