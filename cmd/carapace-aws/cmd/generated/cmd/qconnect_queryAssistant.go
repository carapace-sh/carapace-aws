package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_queryAssistantCmd = &cobra.Command{
	Use:   "query-assistant",
	Short: "This API will be discontinued starting June 1, 2024. To receive generative responses after March 1, 2024, you will need to create a new Assistant in the Amazon Connect console and integrate the Amazon Q in Connect JavaScript library (amazon-q-connectjs) into your applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_queryAssistantCmd).Standalone()

	qconnect_queryAssistantCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_queryAssistantCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_queryAssistantCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_queryAssistantCmd.Flags().String("override-knowledge-base-search-type", "", "The search type to be used against the Knowledge Base for this request.")
	qconnect_queryAssistantCmd.Flags().String("query-condition", "", "Information about how to query content.")
	qconnect_queryAssistantCmd.Flags().String("query-input-data", "", "Information about the query.")
	qconnect_queryAssistantCmd.Flags().String("query-text", "", "The text to search for.")
	qconnect_queryAssistantCmd.Flags().String("session-id", "", "The identifier of the Amazon Q in Connect session.")
	qconnect_queryAssistantCmd.MarkFlagRequired("assistant-id")
	qconnectCmd.AddCommand(qconnect_queryAssistantCmd)
}
