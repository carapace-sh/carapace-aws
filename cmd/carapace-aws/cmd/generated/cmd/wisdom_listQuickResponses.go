package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listQuickResponsesCmd = &cobra.Command{
	Use:   "list-quick-responses",
	Short: "Lists information about quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listQuickResponsesCmd).Standalone()

	wisdom_listQuickResponsesCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_listQuickResponsesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	wisdom_listQuickResponsesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	wisdom_listQuickResponsesCmd.MarkFlagRequired("knowledge-base-id")
	wisdomCmd.AddCommand(wisdom_listQuickResponsesCmd)
}
