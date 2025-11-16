package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listAssessmentTargetsCmd = &cobra.Command{
	Use:   "list-assessment-targets",
	Short: "Lists the ARNs of the assessment targets within this AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listAssessmentTargetsCmd).Standalone()

	inspector_listAssessmentTargetsCmd.Flags().String("filter", "", "You can use this parameter to specify a subset of data to be included in the action's response.")
	inspector_listAssessmentTargetsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	inspector_listAssessmentTargetsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	inspectorCmd.AddCommand(inspector_listAssessmentTargetsCmd)
}
