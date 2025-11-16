package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_searchContentCmd = &cobra.Command{
	Use:   "search-content",
	Short: "Searches for content in a specified knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_searchContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_searchContentCmd).Standalone()

		qconnect_searchContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_searchContentCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		qconnect_searchContentCmd.Flags().String("next-token", "", "The token for the next set of results.")
		qconnect_searchContentCmd.Flags().String("search-expression", "", "The search expression to filter results.")
		qconnect_searchContentCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_searchContentCmd.MarkFlagRequired("search-expression")
	})
	qconnectCmd.AddCommand(qconnect_searchContentCmd)
}
