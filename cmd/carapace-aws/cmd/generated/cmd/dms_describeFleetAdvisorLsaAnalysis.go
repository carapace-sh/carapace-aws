package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeFleetAdvisorLsaAnalysisCmd = &cobra.Command{
	Use:   "describe-fleet-advisor-lsa-analysis",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeFleetAdvisorLsaAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeFleetAdvisorLsaAnalysisCmd).Standalone()

		dms_describeFleetAdvisorLsaAnalysisCmd.Flags().String("max-records", "", "Sets the maximum number of records returned in the response.")
		dms_describeFleetAdvisorLsaAnalysisCmd.Flags().String("next-token", "", "If `NextToken` is returned by a previous response, there are more results available.")
	})
	dmsCmd.AddCommand(dms_describeFleetAdvisorLsaAnalysisCmd)
}
