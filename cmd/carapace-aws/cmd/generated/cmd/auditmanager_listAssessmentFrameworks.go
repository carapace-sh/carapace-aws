package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listAssessmentFrameworksCmd = &cobra.Command{
	Use:   "list-assessment-frameworks",
	Short: "Returns a list of the frameworks that are available in the Audit Manager framework library.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listAssessmentFrameworksCmd).Standalone()

	auditmanager_listAssessmentFrameworksCmd.Flags().String("framework-type", "", "The type of framework, such as a standard framework or a custom framework.")
	auditmanager_listAssessmentFrameworksCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
	auditmanager_listAssessmentFrameworksCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanager_listAssessmentFrameworksCmd.MarkFlagRequired("framework-type")
	auditmanagerCmd.AddCommand(auditmanager_listAssessmentFrameworksCmd)
}
