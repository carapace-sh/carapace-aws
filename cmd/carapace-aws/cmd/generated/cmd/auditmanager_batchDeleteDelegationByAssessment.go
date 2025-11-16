package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_batchDeleteDelegationByAssessmentCmd = &cobra.Command{
	Use:   "batch-delete-delegation-by-assessment",
	Short: "Deletes a batch of delegations for an assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_batchDeleteDelegationByAssessmentCmd).Standalone()

	auditmanager_batchDeleteDelegationByAssessmentCmd.Flags().String("assessment-id", "", "The identifier for the assessment.")
	auditmanager_batchDeleteDelegationByAssessmentCmd.Flags().String("delegation-ids", "", "The identifiers for the delegations.")
	auditmanager_batchDeleteDelegationByAssessmentCmd.MarkFlagRequired("assessment-id")
	auditmanager_batchDeleteDelegationByAssessmentCmd.MarkFlagRequired("delegation-ids")
	auditmanagerCmd.AddCommand(auditmanager_batchDeleteDelegationByAssessmentCmd)
}
