package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_queryAssistantCmd = &cobra.Command{
	Use:   "query-assistant",
	Short: "Performs a manual search against the specified assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_queryAssistantCmd).Standalone()

	wisdom_queryAssistantCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
	wisdom_queryAssistantCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	wisdom_queryAssistantCmd.Flags().String("next-token", "", "The token for the next set of results.")
	wisdom_queryAssistantCmd.Flags().String("query-text", "", "The text to search for.")
	wisdom_queryAssistantCmd.MarkFlagRequired("assistant-id")
	wisdom_queryAssistantCmd.MarkFlagRequired("query-text")
	wisdomCmd.AddCommand(wisdom_queryAssistantCmd)
}
