package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deleteAssessmentFrameworkCmd = &cobra.Command{
	Use:   "delete-assessment-framework",
	Short: "Deletes a custom framework in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deleteAssessmentFrameworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_deleteAssessmentFrameworkCmd).Standalone()

		auditmanager_deleteAssessmentFrameworkCmd.Flags().String("framework-id", "", "The identifier for the custom framework.")
		auditmanager_deleteAssessmentFrameworkCmd.MarkFlagRequired("framework-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_deleteAssessmentFrameworkCmd)
}
