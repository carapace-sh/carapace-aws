package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_prepareQueryCmd = &cobra.Command{
	Use:   "prepare-query",
	Short: "A synchronous operation that allows you to submit a query with parameters to be stored by Timestream for later running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_prepareQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_prepareQueryCmd).Standalone()

		timestreamQuery_prepareQueryCmd.Flags().String("query-string", "", "The Timestream query string that you want to use as a prepared statement.")
		timestreamQuery_prepareQueryCmd.Flags().String("validate-only", "", "By setting this value to `true`, Timestream will only validate that the query string is a valid Timestream query, and not store the prepared query for later use.")
		timestreamQuery_prepareQueryCmd.MarkFlagRequired("query-string")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_prepareQueryCmd)
}
