package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeAppAssessmentCmd = &cobra.Command{
	Use:   "describe-app-assessment",
	Short: "Describes an assessment for an Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeAppAssessmentCmd).Standalone()

	resiliencehub_describeAppAssessmentCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
	resiliencehub_describeAppAssessmentCmd.MarkFlagRequired("assessment-arn")
	resiliencehubCmd.AddCommand(resiliencehub_describeAppAssessmentCmd)
}
