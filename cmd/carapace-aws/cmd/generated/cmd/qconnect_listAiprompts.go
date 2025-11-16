package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAipromptsCmd = &cobra.Command{
	Use:   "list-aiprompts",
	Short: "Lists the AI Prompts available on the Amazon Q in Connect assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAipromptsCmd).Standalone()

	qconnect_listAipromptsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listAipromptsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAipromptsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listAipromptsCmd.Flags().String("origin", "", "The origin of the AI Prompts to be listed.")
	qconnect_listAipromptsCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_listAipromptsCmd)
}
