package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_deleteAssessmentTemplateCmd = &cobra.Command{
	Use:   "delete-assessment-template",
	Short: "Deletes the assessment template that is specified by the ARN of the assessment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_deleteAssessmentTemplateCmd).Standalone()

	inspector_deleteAssessmentTemplateCmd.Flags().String("assessment-template-arn", "", "The ARN that specifies the assessment template that you want to delete.")
	inspector_deleteAssessmentTemplateCmd.MarkFlagRequired("assessment-template-arn")
	inspectorCmd.AddCommand(inspector_deleteAssessmentTemplateCmd)
}
