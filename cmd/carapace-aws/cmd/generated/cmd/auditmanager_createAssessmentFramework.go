package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_createAssessmentFrameworkCmd = &cobra.Command{
	Use:   "create-assessment-framework",
	Short: "Creates a custom framework in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_createAssessmentFrameworkCmd).Standalone()

	auditmanager_createAssessmentFrameworkCmd.Flags().String("compliance-type", "", "The compliance type that the new custom framework supports, such as CIS or HIPAA.")
	auditmanager_createAssessmentFrameworkCmd.Flags().String("control-sets", "", "The control sets that are associated with the framework.")
	auditmanager_createAssessmentFrameworkCmd.Flags().String("description", "", "An optional description for the new custom framework.")
	auditmanager_createAssessmentFrameworkCmd.Flags().String("name", "", "The name of the new custom framework.")
	auditmanager_createAssessmentFrameworkCmd.Flags().String("tags", "", "The tags that are associated with the framework.")
	auditmanager_createAssessmentFrameworkCmd.MarkFlagRequired("control-sets")
	auditmanager_createAssessmentFrameworkCmd.MarkFlagRequired("name")
	auditmanagerCmd.AddCommand(auditmanager_createAssessmentFrameworkCmd)
}
