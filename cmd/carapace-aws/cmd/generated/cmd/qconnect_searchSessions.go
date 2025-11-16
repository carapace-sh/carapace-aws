package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_searchSessionsCmd = &cobra.Command{
	Use:   "search-sessions",
	Short: "Searches for sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_searchSessionsCmd).Standalone()

	qconnect_searchSessionsCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_searchSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_searchSessionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_searchSessionsCmd.Flags().String("search-expression", "", "The search expression to filter results.")
	qconnect_searchSessionsCmd.MarkFlagRequired("assistant-id")
	qconnect_searchSessionsCmd.MarkFlagRequired("search-expression")
	qconnectCmd.AddCommand(qconnect_searchSessionsCmd)
}
