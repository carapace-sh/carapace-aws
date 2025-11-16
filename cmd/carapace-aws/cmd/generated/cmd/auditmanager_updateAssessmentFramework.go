package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateAssessmentFrameworkCmd = &cobra.Command{
	Use:   "update-assessment-framework",
	Short: "Updates a custom framework in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateAssessmentFrameworkCmd).Standalone()

	auditmanager_updateAssessmentFrameworkCmd.Flags().String("compliance-type", "", "The compliance type that the new custom framework supports, such as CIS or HIPAA.")
	auditmanager_updateAssessmentFrameworkCmd.Flags().String("control-sets", "", "The control sets that are associated with the framework.")
	auditmanager_updateAssessmentFrameworkCmd.Flags().String("description", "", "The description of the updated framework.")
	auditmanager_updateAssessmentFrameworkCmd.Flags().String("framework-id", "", "The unique identifier for the framework.")
	auditmanager_updateAssessmentFrameworkCmd.Flags().String("name", "", "The name of the framework to be updated.")
	auditmanager_updateAssessmentFrameworkCmd.MarkFlagRequired("control-sets")
	auditmanager_updateAssessmentFrameworkCmd.MarkFlagRequired("framework-id")
	auditmanager_updateAssessmentFrameworkCmd.MarkFlagRequired("name")
	auditmanagerCmd.AddCommand(auditmanager_updateAssessmentFrameworkCmd)
}
