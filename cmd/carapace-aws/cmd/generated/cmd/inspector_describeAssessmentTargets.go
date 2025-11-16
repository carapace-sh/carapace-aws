package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeAssessmentTargetsCmd = &cobra.Command{
	Use:   "describe-assessment-targets",
	Short: "Describes the assessment targets that are specified by the ARNs of the assessment targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeAssessmentTargetsCmd).Standalone()

	inspector_describeAssessmentTargetsCmd.Flags().String("assessment-target-arns", "", "The ARNs that specifies the assessment targets that you want to describe.")
	inspector_describeAssessmentTargetsCmd.MarkFlagRequired("assessment-target-arns")
	inspectorCmd.AddCommand(inspector_describeAssessmentTargetsCmd)
}
