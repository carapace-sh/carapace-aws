package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbrecommendationsCmd = &cobra.Command{
	Use:   "describe-dbrecommendations",
	Short: "Describes the recommendations to resolve the issues for your DB instances, DB clusters, and DB parameter groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbrecommendationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbrecommendationsCmd).Standalone()

		rds_describeDbrecommendationsCmd.Flags().String("filters", "", "A filter that specifies one or more recommendations to describe.")
		rds_describeDbrecommendationsCmd.Flags().String("last-updated-after", "", "A filter to include only the recommendations that were updated after this specified time.")
		rds_describeDbrecommendationsCmd.Flags().String("last-updated-before", "", "A filter to include only the recommendations that were updated before this specified time.")
		rds_describeDbrecommendationsCmd.Flags().String("locale", "", "The language that you choose to return the list of recommendations.")
		rds_describeDbrecommendationsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBRecommendations` request.")
		rds_describeDbrecommendationsCmd.Flags().String("max-records", "", "The maximum number of recommendations to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbrecommendationsCmd)
}
