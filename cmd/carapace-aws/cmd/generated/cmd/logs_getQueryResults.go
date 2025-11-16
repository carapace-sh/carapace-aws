package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getQueryResultsCmd = &cobra.Command{
	Use:   "get-query-results",
	Short: "Returns the results from the specified query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getQueryResultsCmd).Standalone()

	logs_getQueryResultsCmd.Flags().String("query-id", "", "The ID number of the query.")
	logs_getQueryResultsCmd.MarkFlagRequired("query-id")
	logsCmd.AddCommand(logs_getQueryResultsCmd)
}
