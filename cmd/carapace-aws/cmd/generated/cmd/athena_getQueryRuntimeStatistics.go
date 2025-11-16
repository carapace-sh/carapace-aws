package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getQueryRuntimeStatisticsCmd = &cobra.Command{
	Use:   "get-query-runtime-statistics",
	Short: "Returns query execution runtime statistics related to a single execution of a query if you have access to the workgroup in which the query ran.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getQueryRuntimeStatisticsCmd).Standalone()

	athena_getQueryRuntimeStatisticsCmd.Flags().String("query-execution-id", "", "The unique ID of the query execution.")
	athena_getQueryRuntimeStatisticsCmd.MarkFlagRequired("query-execution-id")
	athenaCmd.AddCommand(athena_getQueryRuntimeStatisticsCmd)
}
