package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listAssessmentsCmd = &cobra.Command{
	Use:   "list-assessments",
	Short: "Returns a list of current and past assessments from Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listAssessmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listAssessmentsCmd).Standalone()

		auditmanager_listAssessmentsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_listAssessmentsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_listAssessmentsCmd.Flags().String("status", "", "The current status of the assessment.")
	})
	auditmanagerCmd.AddCommand(auditmanager_listAssessmentsCmd)
}
