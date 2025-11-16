package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "`Query` is a synchronous operation that enables you to run a query against your Amazon Timestream data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_queryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_queryCmd).Standalone()

		timestreamQuery_queryCmd.Flags().String("client-token", "", "Unique, case-sensitive string of up to 64 ASCII characters specified when a `Query` request is made.")
		timestreamQuery_queryCmd.Flags().String("max-rows", "", "The total number of rows to be returned in the `Query` output.")
		timestreamQuery_queryCmd.Flags().String("next-token", "", "A pagination token used to return a set of results.")
		timestreamQuery_queryCmd.Flags().String("query-insights", "", "Encapsulates settings for enabling `QueryInsights`.")
		timestreamQuery_queryCmd.Flags().String("query-string", "", "The query to be run by Timestream.")
		timestreamQuery_queryCmd.MarkFlagRequired("query-string")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_queryCmd)
}
