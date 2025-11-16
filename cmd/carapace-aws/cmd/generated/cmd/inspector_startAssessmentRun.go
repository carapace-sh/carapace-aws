package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_startAssessmentRunCmd = &cobra.Command{
	Use:   "start-assessment-run",
	Short: "Starts the assessment run specified by the ARN of the assessment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_startAssessmentRunCmd).Standalone()

	inspector_startAssessmentRunCmd.Flags().String("assessment-run-name", "", "You can specify the name for the assessment run.")
	inspector_startAssessmentRunCmd.Flags().String("assessment-template-arn", "", "The ARN of the assessment template of the assessment run that you want to start.")
	inspector_startAssessmentRunCmd.MarkFlagRequired("assessment-template-arn")
	inspectorCmd.AddCommand(inspector_startAssessmentRunCmd)
}
