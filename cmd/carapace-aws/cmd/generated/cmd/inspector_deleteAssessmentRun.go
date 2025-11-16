package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_deleteAssessmentRunCmd = &cobra.Command{
	Use:   "delete-assessment-run",
	Short: "Deletes the assessment run that is specified by the ARN of the assessment run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_deleteAssessmentRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_deleteAssessmentRunCmd).Standalone()

		inspector_deleteAssessmentRunCmd.Flags().String("assessment-run-arn", "", "The ARN that specifies the assessment run that you want to delete.")
		inspector_deleteAssessmentRunCmd.MarkFlagRequired("assessment-run-arn")
	})
	inspectorCmd.AddCommand(inspector_deleteAssessmentRunCmd)
}
