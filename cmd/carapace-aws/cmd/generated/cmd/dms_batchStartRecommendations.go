package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_batchStartRecommendationsCmd = &cobra.Command{
	Use:   "batch-start-recommendations",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_batchStartRecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_batchStartRecommendationsCmd).Standalone()

		dms_batchStartRecommendationsCmd.Flags().String("data", "", "Provides information about source databases to analyze.")
	})
	dmsCmd.AddCommand(dms_batchStartRecommendationsCmd)
}
