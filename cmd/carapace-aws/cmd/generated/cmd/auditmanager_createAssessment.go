package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_createAssessmentCmd = &cobra.Command{
	Use:   "create-assessment",
	Short: "Creates an assessment in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_createAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_createAssessmentCmd).Standalone()

		auditmanager_createAssessmentCmd.Flags().String("assessment-reports-destination", "", "The assessment report storage destination for the assessment that's being created.")
		auditmanager_createAssessmentCmd.Flags().String("description", "", "The optional description of the assessment to be created.")
		auditmanager_createAssessmentCmd.Flags().String("framework-id", "", "The identifier for the framework that the assessment will be created from.")
		auditmanager_createAssessmentCmd.Flags().String("name", "", "The name of the assessment to be created.")
		auditmanager_createAssessmentCmd.Flags().String("roles", "", "The list of roles for the assessment.")
		auditmanager_createAssessmentCmd.Flags().String("scope", "", "")
		auditmanager_createAssessmentCmd.Flags().String("tags", "", "The tags that are associated with the assessment.")
		auditmanager_createAssessmentCmd.MarkFlagRequired("assessment-reports-destination")
		auditmanager_createAssessmentCmd.MarkFlagRequired("framework-id")
		auditmanager_createAssessmentCmd.MarkFlagRequired("name")
		auditmanager_createAssessmentCmd.MarkFlagRequired("roles")
		auditmanager_createAssessmentCmd.MarkFlagRequired("scope")
	})
	auditmanagerCmd.AddCommand(auditmanager_createAssessmentCmd)
}
