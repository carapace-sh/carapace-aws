package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppAssessmentResourceDriftsCmd = &cobra.Command{
	Use:   "list-app-assessment-resource-drifts",
	Short: "List of resource drifts that were detected while running an assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppAssessmentResourceDriftsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listAppAssessmentResourceDriftsCmd).Standalone()

		resiliencehub_listAppAssessmentResourceDriftsCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
		resiliencehub_listAppAssessmentResourceDriftsCmd.Flags().String("max-results", "", "Maximum number of drift results to include in the response.")
		resiliencehub_listAppAssessmentResourceDriftsCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
		resiliencehub_listAppAssessmentResourceDriftsCmd.MarkFlagRequired("assessment-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listAppAssessmentResourceDriftsCmd)
}
