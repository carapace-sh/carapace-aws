package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_stopAssessmentCmd = &cobra.Command{
	Use:   "stop-assessment",
	Short: "Stops the assessment of an on-premises environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_stopAssessmentCmd).Standalone()

	migrationhubstrategy_stopAssessmentCmd.Flags().String("assessment-id", "", "The `assessmentId` returned by [StartAssessment]().")
	migrationhubstrategy_stopAssessmentCmd.MarkFlagRequired("assessment-id")
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_stopAssessmentCmd)
}
