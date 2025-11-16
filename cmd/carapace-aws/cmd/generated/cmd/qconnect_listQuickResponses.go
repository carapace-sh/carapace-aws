package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listQuickResponsesCmd = &cobra.Command{
	Use:   "list-quick-responses",
	Short: "Lists information about quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listQuickResponsesCmd).Standalone()

	qconnect_listQuickResponsesCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_listQuickResponsesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listQuickResponsesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listQuickResponsesCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_listQuickResponsesCmd)
}
