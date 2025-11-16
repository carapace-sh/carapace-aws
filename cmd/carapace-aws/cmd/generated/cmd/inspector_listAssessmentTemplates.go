package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listAssessmentTemplatesCmd = &cobra.Command{
	Use:   "list-assessment-templates",
	Short: "Lists the assessment templates that correspond to the assessment targets that are specified by the ARNs of the assessment targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listAssessmentTemplatesCmd).Standalone()

	inspector_listAssessmentTemplatesCmd.Flags().String("assessment-target-arns", "", "A list of ARNs that specifies the assessment targets whose assessment templates you want to list.")
	inspector_listAssessmentTemplatesCmd.Flags().String("filter", "", "You can use this parameter to specify a subset of data to be included in the action's response.")
	inspector_listAssessmentTemplatesCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	inspector_listAssessmentTemplatesCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	inspectorCmd.AddCommand(inspector_listAssessmentTemplatesCmd)
}
