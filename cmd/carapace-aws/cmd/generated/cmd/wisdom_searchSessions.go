package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_searchSessionsCmd = &cobra.Command{
	Use:   "search-sessions",
	Short: "Searches for sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_searchSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_searchSessionsCmd).Standalone()

		wisdom_searchSessionsCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
		wisdom_searchSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		wisdom_searchSessionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		wisdom_searchSessionsCmd.Flags().String("search-expression", "", "The search expression to filter results.")
		wisdom_searchSessionsCmd.MarkFlagRequired("assistant-id")
		wisdom_searchSessionsCmd.MarkFlagRequired("search-expression")
	})
	wisdomCmd.AddCommand(wisdom_searchSessionsCmd)
}
