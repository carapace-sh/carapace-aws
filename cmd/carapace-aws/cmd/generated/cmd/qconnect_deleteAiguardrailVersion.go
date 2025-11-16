package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteAiguardrailVersionCmd = &cobra.Command{
	Use:   "delete-aiguardrail-version",
	Short: "Delete and Amazon Q in Connect AI Guardrail version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteAiguardrailVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteAiguardrailVersionCmd).Standalone()

		qconnect_deleteAiguardrailVersionCmd.Flags().String("ai-guardrail-id", "", "The identifier of the Amazon Q in Connect AI Guardrail.")
		qconnect_deleteAiguardrailVersionCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_deleteAiguardrailVersionCmd.Flags().String("version-number", "", "The version number of the AI Guardrail version to be deleted.")
		qconnect_deleteAiguardrailVersionCmd.MarkFlagRequired("ai-guardrail-id")
		qconnect_deleteAiguardrailVersionCmd.MarkFlagRequired("assistant-id")
		qconnect_deleteAiguardrailVersionCmd.MarkFlagRequired("version-number")
	})
	qconnectCmd.AddCommand(qconnect_deleteAiguardrailVersionCmd)
}
