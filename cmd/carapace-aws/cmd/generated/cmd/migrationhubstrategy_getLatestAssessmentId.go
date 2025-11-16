package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_getLatestAssessmentIdCmd = &cobra.Command{
	Use:   "get-latest-assessment-id",
	Short: "Retrieve the latest ID of a specific assessment task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_getLatestAssessmentIdCmd).Standalone()

	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_getLatestAssessmentIdCmd)
}
