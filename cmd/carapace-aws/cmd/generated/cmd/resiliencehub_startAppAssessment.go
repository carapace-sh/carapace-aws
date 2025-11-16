package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_startAppAssessmentCmd = &cobra.Command{
	Use:   "start-app-assessment",
	Short: "Creates a new application assessment for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_startAppAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_startAppAssessmentCmd).Standalone()

		resiliencehub_startAppAssessmentCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
		resiliencehub_startAppAssessmentCmd.Flags().String("app-version", "", "The version of the application.")
		resiliencehub_startAppAssessmentCmd.Flags().String("assessment-name", "", "The name for the assessment.")
		resiliencehub_startAppAssessmentCmd.Flags().String("client-token", "", "Used for an idempotency token.")
		resiliencehub_startAppAssessmentCmd.Flags().String("tags", "", "Tags assigned to the resource.")
		resiliencehub_startAppAssessmentCmd.MarkFlagRequired("app-arn")
		resiliencehub_startAppAssessmentCmd.MarkFlagRequired("app-version")
		resiliencehub_startAppAssessmentCmd.MarkFlagRequired("assessment-name")
	})
	resiliencehubCmd.AddCommand(resiliencehub_startAppAssessmentCmd)
}
