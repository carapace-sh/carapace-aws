package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeAssessmentTemplatesCmd = &cobra.Command{
	Use:   "describe-assessment-templates",
	Short: "Describes the assessment templates that are specified by the ARNs of the assessment templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeAssessmentTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_describeAssessmentTemplatesCmd).Standalone()

		inspector_describeAssessmentTemplatesCmd.Flags().String("assessment-template-arns", "", "")
		inspector_describeAssessmentTemplatesCmd.MarkFlagRequired("assessment-template-arns")
	})
	inspectorCmd.AddCommand(inspector_describeAssessmentTemplatesCmd)
}
