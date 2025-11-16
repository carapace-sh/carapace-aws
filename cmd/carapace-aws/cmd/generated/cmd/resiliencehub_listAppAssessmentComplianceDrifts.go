package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppAssessmentComplianceDriftsCmd = &cobra.Command{
	Use:   "list-app-assessment-compliance-drifts",
	Short: "List of compliance drifts that were detected while running an assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppAssessmentComplianceDriftsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listAppAssessmentComplianceDriftsCmd).Standalone()

		resiliencehub_listAppAssessmentComplianceDriftsCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
		resiliencehub_listAppAssessmentComplianceDriftsCmd.Flags().String("max-results", "", "Maximum number of compliance drifts requested.")
		resiliencehub_listAppAssessmentComplianceDriftsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listAppAssessmentComplianceDriftsCmd.MarkFlagRequired("assessment-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listAppAssessmentComplianceDriftsCmd)
}
