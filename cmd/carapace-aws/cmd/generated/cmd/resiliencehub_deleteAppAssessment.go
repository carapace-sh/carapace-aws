package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_deleteAppAssessmentCmd = &cobra.Command{
	Use:   "delete-app-assessment",
	Short: "Deletes an Resilience Hub application assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_deleteAppAssessmentCmd).Standalone()

	resiliencehub_deleteAppAssessmentCmd.Flags().String("assessment-arn", "", "Amazon Resource Name (ARN) of the assessment.")
	resiliencehub_deleteAppAssessmentCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehub_deleteAppAssessmentCmd.MarkFlagRequired("assessment-arn")
	resiliencehubCmd.AddCommand(resiliencehub_deleteAppAssessmentCmd)
}
