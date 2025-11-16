package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getAssessmentCmd = &cobra.Command{
	Use:   "get-assessment",
	Short: "Retrieves the status of an on-going assessment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getAssessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_getAssessmentCmd).Standalone()

		migrationhubstrategy_getAssessmentCmd.Flags().String("id", "", "The `assessmentid` returned by [StartAssessment]().")
		migrationhubstrategy_getAssessmentCmd.MarkFlagRequired("id")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getAssessmentCmd)
}
