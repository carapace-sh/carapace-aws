package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeFleetAdvisorSchemaObjectSummaryCmd = &cobra.Command{
	Use:   "describe-fleet-advisor-schema-object-summary",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeFleetAdvisorSchemaObjectSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeFleetAdvisorSchemaObjectSummaryCmd).Standalone()

		dms_describeFleetAdvisorSchemaObjectSummaryCmd.Flags().String("filters", "", "If you specify any of the following filters, the output includes information for only those schema objects that meet the filter criteria:")
		dms_describeFleetAdvisorSchemaObjectSummaryCmd.Flags().String("max-records", "", "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.")
		dms_describeFleetAdvisorSchemaObjectSummaryCmd.Flags().String("next-token", "", "If `NextToken` is returned by a previous response, there are more results available.")
	})
	dmsCmd.AddCommand(dms_describeFleetAdvisorSchemaObjectSummaryCmd)
}
