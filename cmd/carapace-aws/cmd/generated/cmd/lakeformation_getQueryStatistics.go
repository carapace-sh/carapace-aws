package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getQueryStatisticsCmd = &cobra.Command{
	Use:   "get-query-statistics",
	Short: "Retrieves statistics on the planning and execution of a query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getQueryStatisticsCmd).Standalone()

	lakeformation_getQueryStatisticsCmd.Flags().String("query-id", "", "The ID of the plan query operation.")
	lakeformation_getQueryStatisticsCmd.MarkFlagRequired("query-id")
	lakeformationCmd.AddCommand(lakeformation_getQueryStatisticsCmd)
}
