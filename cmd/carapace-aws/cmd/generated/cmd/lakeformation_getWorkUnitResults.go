package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getWorkUnitResultsCmd = &cobra.Command{
	Use:   "get-work-unit-results",
	Short: "Returns the work units resulting from the query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getWorkUnitResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getWorkUnitResultsCmd).Standalone()

		lakeformation_getWorkUnitResultsCmd.Flags().String("query-id", "", "The ID of the plan query operation for which to get results.")
		lakeformation_getWorkUnitResultsCmd.Flags().String("work-unit-id", "", "The work unit ID for which to get results.")
		lakeformation_getWorkUnitResultsCmd.Flags().String("work-unit-token", "", "A work token used to query the execution service.")
		lakeformation_getWorkUnitResultsCmd.MarkFlagRequired("query-id")
		lakeformation_getWorkUnitResultsCmd.MarkFlagRequired("work-unit-id")
		lakeformation_getWorkUnitResultsCmd.MarkFlagRequired("work-unit-token")
	})
	lakeformationCmd.AddCommand(lakeformation_getWorkUnitResultsCmd)
}
