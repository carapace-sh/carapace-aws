package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAipromptVersionsCmd = &cobra.Command{
	Use:   "list-aiprompt-versions",
	Short: "Lists AI Prompt versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAipromptVersionsCmd).Standalone()

	qconnect_listAipromptVersionsCmd.Flags().String("ai-prompt-id", "", "The identifier of the Amazon Q in Connect AI prompt for which versions are to be listed.")
	qconnect_listAipromptVersionsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_listAipromptVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAipromptVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listAipromptVersionsCmd.Flags().String("origin", "", "The origin of the AI Prompt versions to be listed.")
	qconnect_listAipromptVersionsCmd.MarkFlagRequired("ai-prompt-id")
	qconnect_listAipromptVersionsCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_listAipromptVersionsCmd)
}
