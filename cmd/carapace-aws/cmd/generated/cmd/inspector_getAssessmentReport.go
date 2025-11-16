package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_getAssessmentReportCmd = &cobra.Command{
	Use:   "get-assessment-report",
	Short: "Produces an assessment report that includes detailed and comprehensive results of a specified assessment run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_getAssessmentReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_getAssessmentReportCmd).Standalone()

		inspector_getAssessmentReportCmd.Flags().String("assessment-run-arn", "", "The ARN that specifies the assessment run for which you want to generate a report.")
		inspector_getAssessmentReportCmd.Flags().String("report-file-format", "", "Specifies the file format (html or pdf) of the assessment report that you want to generate.")
		inspector_getAssessmentReportCmd.Flags().String("report-type", "", "Specifies the type of the assessment report that you want to generate.")
		inspector_getAssessmentReportCmd.MarkFlagRequired("assessment-run-arn")
		inspector_getAssessmentReportCmd.MarkFlagRequired("report-file-format")
		inspector_getAssessmentReportCmd.MarkFlagRequired("report-type")
	})
	inspectorCmd.AddCommand(inspector_getAssessmentReportCmd)
}
