package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_deleteAssessmentTargetCmd = &cobra.Command{
	Use:   "delete-assessment-target",
	Short: "Deletes the assessment target that is specified by the ARN of the assessment target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_deleteAssessmentTargetCmd).Standalone()

	inspector_deleteAssessmentTargetCmd.Flags().String("assessment-target-arn", "", "The ARN that specifies the assessment target that you want to delete.")
	inspector_deleteAssessmentTargetCmd.MarkFlagRequired("assessment-target-arn")
	inspectorCmd.AddCommand(inspector_deleteAssessmentTargetCmd)
}
