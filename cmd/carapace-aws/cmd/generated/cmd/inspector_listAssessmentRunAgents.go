package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_listAssessmentRunAgentsCmd = &cobra.Command{
	Use:   "list-assessment-run-agents",
	Short: "Lists the agents of the assessment runs that are specified by the ARNs of the assessment runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_listAssessmentRunAgentsCmd).Standalone()

	inspector_listAssessmentRunAgentsCmd.Flags().String("assessment-run-arn", "", "The ARN that specifies the assessment run whose agents you want to list.")
	inspector_listAssessmentRunAgentsCmd.Flags().String("filter", "", "You can use this parameter to specify a subset of data to be included in the action's response.")
	inspector_listAssessmentRunAgentsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
	inspector_listAssessmentRunAgentsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	inspector_listAssessmentRunAgentsCmd.MarkFlagRequired("assessment-run-arn")
	inspectorCmd.AddCommand(inspector_listAssessmentRunAgentsCmd)
}
