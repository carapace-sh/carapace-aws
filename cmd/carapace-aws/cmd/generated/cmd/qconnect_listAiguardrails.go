package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAiguardrailsCmd = &cobra.Command{
	Use:   "list-aiguardrails",
	Short: "Lists the AI Guardrails available on the Amazon Q in Connect assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAiguardrailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_listAiguardrailsCmd).Standalone()

		qconnect_listAiguardrailsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
		qconnect_listAiguardrailsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		qconnect_listAiguardrailsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		qconnect_listAiguardrailsCmd.MarkFlagRequired("assistant-id")
	})
	qconnectCmd.AddCommand(qconnect_listAiguardrailsCmd)
}
