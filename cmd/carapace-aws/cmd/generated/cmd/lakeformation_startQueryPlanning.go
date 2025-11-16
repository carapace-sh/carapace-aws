package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_startQueryPlanningCmd = &cobra.Command{
	Use:   "start-query-planning",
	Short: "Submits a request to process a query statement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_startQueryPlanningCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_startQueryPlanningCmd).Standalone()

		lakeformation_startQueryPlanningCmd.Flags().String("query-planning-context", "", "A structure containing information about the query plan.")
		lakeformation_startQueryPlanningCmd.Flags().String("query-string", "", "A PartiQL query statement used as an input to the planner service.")
		lakeformation_startQueryPlanningCmd.MarkFlagRequired("query-planning-context")
		lakeformation_startQueryPlanningCmd.MarkFlagRequired("query-string")
	})
	lakeformationCmd.AddCommand(lakeformation_startQueryPlanningCmd)
}
