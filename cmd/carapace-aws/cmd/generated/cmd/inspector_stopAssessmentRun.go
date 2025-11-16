package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_stopAssessmentRunCmd = &cobra.Command{
	Use:   "stop-assessment-run",
	Short: "Stops the assessment run that is specified by the ARN of the assessment run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_stopAssessmentRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_stopAssessmentRunCmd).Standalone()

		inspector_stopAssessmentRunCmd.Flags().String("assessment-run-arn", "", "The ARN of the assessment run that you want to stop.")
		inspector_stopAssessmentRunCmd.Flags().String("stop-action", "", "An input option that can be set to either START\\_EVALUATION or SKIP\\_EVALUATION.")
		inspector_stopAssessmentRunCmd.MarkFlagRequired("assessment-run-arn")
	})
	inspectorCmd.AddCommand(inspector_stopAssessmentRunCmd)
}
