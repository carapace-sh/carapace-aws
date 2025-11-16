package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_batchCreateDelegationByAssessmentCmd = &cobra.Command{
	Use:   "batch-create-delegation-by-assessment",
	Short: "Creates a batch of delegations for an assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_batchCreateDelegationByAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_batchCreateDelegationByAssessmentCmd).Standalone()

		auditmanager_batchCreateDelegationByAssessmentCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
		auditmanager_batchCreateDelegationByAssessmentCmd.Flags().String("create-delegation-requests", "", "The API request to batch create delegations in Audit Manager.")
		auditmanager_batchCreateDelegationByAssessmentCmd.MarkFlagRequired("assessment-id")
		auditmanager_batchCreateDelegationByAssessmentCmd.MarkFlagRequired("create-delegation-requests")
	})
	auditmanagerCmd.AddCommand(auditmanager_batchCreateDelegationByAssessmentCmd)
}
