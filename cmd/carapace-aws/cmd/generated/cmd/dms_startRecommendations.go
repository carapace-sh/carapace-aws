package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startRecommendationsCmd = &cobra.Command{
	Use:   "start-recommendations",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startRecommendationsCmd).Standalone()

	dms_startRecommendationsCmd.Flags().String("database-id", "", "The identifier of the source database to analyze and provide recommendations for.")
	dms_startRecommendationsCmd.Flags().String("settings", "", "The settings in JSON format that Fleet Advisor uses to determine target engine recommendations.")
	dms_startRecommendationsCmd.MarkFlagRequired("database-id")
	dms_startRecommendationsCmd.MarkFlagRequired("settings")
	dmsCmd.AddCommand(dms_startRecommendationsCmd)
}
