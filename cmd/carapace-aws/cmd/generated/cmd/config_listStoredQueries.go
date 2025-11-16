package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listStoredQueriesCmd = &cobra.Command{
	Use:   "list-stored-queries",
	Short: "Lists the stored queries for a single Amazon Web Services account and a single Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listStoredQueriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_listStoredQueriesCmd).Standalone()

		config_listStoredQueriesCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		config_listStoredQueriesCmd.Flags().String("next-token", "", "The nextToken string returned in a previous request that you use to request the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_listStoredQueriesCmd)
}
