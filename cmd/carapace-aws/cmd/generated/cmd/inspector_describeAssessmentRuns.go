package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeAssessmentRunsCmd = &cobra.Command{
	Use:   "describe-assessment-runs",
	Short: "Describes the assessment runs that are specified by the ARNs of the assessment runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeAssessmentRunsCmd).Standalone()

	inspector_describeAssessmentRunsCmd.Flags().String("assessment-run-arns", "", "The ARN that specifies the assessment run that you want to describe.")
	inspector_describeAssessmentRunsCmd.MarkFlagRequired("assessment-run-arns")
	inspectorCmd.AddCommand(inspector_describeAssessmentRunsCmd)
}
