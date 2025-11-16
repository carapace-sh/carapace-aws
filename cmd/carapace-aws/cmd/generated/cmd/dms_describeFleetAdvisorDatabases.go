package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeFleetAdvisorDatabasesCmd = &cobra.Command{
	Use:   "describe-fleet-advisor-databases",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeFleetAdvisorDatabasesCmd).Standalone()

	dms_describeFleetAdvisorDatabasesCmd.Flags().String("filters", "", "If you specify any of the following filters, the output includes information for only those databases that meet the filter criteria:")
	dms_describeFleetAdvisorDatabasesCmd.Flags().String("max-records", "", "Sets the maximum number of records returned in the response.")
	dms_describeFleetAdvisorDatabasesCmd.Flags().String("next-token", "", "If `NextToken` is returned by a previous response, there are more results available.")
	dmsCmd.AddCommand(dms_describeFleetAdvisorDatabasesCmd)
}
