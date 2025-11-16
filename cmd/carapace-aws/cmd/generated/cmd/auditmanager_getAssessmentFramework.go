package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getAssessmentFrameworkCmd = &cobra.Command{
	Use:   "get-assessment-framework",
	Short: "Gets information about a specified framework.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getAssessmentFrameworkCmd).Standalone()

	auditmanager_getAssessmentFrameworkCmd.Flags().String("framework-id", "", "The identifier for the framework.")
	auditmanager_getAssessmentFrameworkCmd.MarkFlagRequired("framework-id")
	auditmanagerCmd.AddCommand(auditmanager_getAssessmentFrameworkCmd)
}
