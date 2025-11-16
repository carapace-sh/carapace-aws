package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAiguardrailVersionsCmd = &cobra.Command{
	Use:   "list-aiguardrail-versions",
	Short: "Lists AI Guardrail versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAiguardrailVersionsCmd).Standalone()

	qconnect_listAiguardrailVersionsCmd.Flags().String("ai-guardrail-id", "", "The identifier of the Amazon Q in Connect AI Guardrail for which versions are to be listed.")
	qconnect_listAiguardrailVersionsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listAiguardrailVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAiguardrailVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listAiguardrailVersionsCmd.MarkFlagRequired("ai-guardrail-id")
	qconnect_listAiguardrailVersionsCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_listAiguardrailVersionsCmd)
}
