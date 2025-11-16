package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listAssessmentFrameworkShareRequestsCmd = &cobra.Command{
	Use:   "list-assessment-framework-share-requests",
	Short: "Returns a list of sent or received share requests for custom frameworks in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listAssessmentFrameworkShareRequestsCmd).Standalone()

	auditmanager_listAssessmentFrameworkShareRequestsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
	auditmanager_listAssessmentFrameworkShareRequestsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	auditmanager_listAssessmentFrameworkShareRequestsCmd.Flags().String("request-type", "", "Specifies whether the share request is a sent request or a received request.")
	auditmanager_listAssessmentFrameworkShareRequestsCmd.MarkFlagRequired("request-type")
	auditmanagerCmd.AddCommand(auditmanager_listAssessmentFrameworkShareRequestsCmd)
}
