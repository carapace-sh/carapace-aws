package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_updateAssessmentTargetCmd = &cobra.Command{
	Use:   "update-assessment-target",
	Short: "Updates the assessment target that is specified by the ARN of the assessment target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_updateAssessmentTargetCmd).Standalone()

	inspector_updateAssessmentTargetCmd.Flags().String("assessment-target-arn", "", "The ARN of the assessment target that you want to update.")
	inspector_updateAssessmentTargetCmd.Flags().String("assessment-target-name", "", "The name of the assessment target that you want to update.")
	inspector_updateAssessmentTargetCmd.Flags().String("resource-group-arn", "", "The ARN of the resource group that is used to specify the new resource group to associate with the assessment target.")
	inspector_updateAssessmentTargetCmd.MarkFlagRequired("assessment-target-arn")
	inspector_updateAssessmentTargetCmd.MarkFlagRequired("assessment-target-name")
	inspectorCmd.AddCommand(inspector_updateAssessmentTargetCmd)
}
