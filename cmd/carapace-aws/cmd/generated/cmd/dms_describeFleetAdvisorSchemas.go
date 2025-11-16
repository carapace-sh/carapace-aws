package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeFleetAdvisorSchemasCmd = &cobra.Command{
	Use:   "describe-fleet-advisor-schemas",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeFleetAdvisorSchemasCmd).Standalone()

	dms_describeFleetAdvisorSchemasCmd.Flags().String("filters", "", "If you specify any of the following filters, the output includes information for only those schemas that meet the filter criteria:")
	dms_describeFleetAdvisorSchemasCmd.Flags().String("max-records", "", "Sets the maximum number of records returned in the response.")
	dms_describeFleetAdvisorSchemasCmd.Flags().String("next-token", "", "If `NextToken` is returned by a previous response, there are more results available.")
	dmsCmd.AddCommand(dms_describeFleetAdvisorSchemasCmd)
}
