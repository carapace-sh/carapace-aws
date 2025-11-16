package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listAssessmentReportsCmd = &cobra.Command{
	Use:   "list-assessment-reports",
	Short: "Returns a list of assessment reports created in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listAssessmentReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listAssessmentReportsCmd).Standalone()

		auditmanager_listAssessmentReportsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_listAssessmentReportsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	})
	auditmanagerCmd.AddCommand(auditmanager_listAssessmentReportsCmd)
}
