package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listAssessmentRunsCmd = &cobra.Command{
	Use:   "list-assessment-runs",
	Short: "Lists the assessment runs that correspond to the assessment templates that are specified by the ARNs of the assessment templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listAssessmentRunsCmd).Standalone()

	inspector_listAssessmentRunsCmd.Flags().String("assessment-template-arns", "", "The ARNs that specify the assessment templates whose assessment runs you want to list.")
	inspector_listAssessmentRunsCmd.Flags().String("filter", "", "You can use this parameter to specify a subset of data to be included in the action's response.")
	inspector_listAssessmentRunsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
	inspector_listAssessmentRunsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	inspectorCmd.AddCommand(inspector_listAssessmentRunsCmd)
}
