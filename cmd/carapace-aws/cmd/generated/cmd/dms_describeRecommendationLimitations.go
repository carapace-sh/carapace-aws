package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeRecommendationLimitationsCmd = &cobra.Command{
	Use:   "describe-recommendation-limitations",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeRecommendationLimitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeRecommendationLimitationsCmd).Standalone()

		dms_describeRecommendationLimitationsCmd.Flags().String("filters", "", "Filters applied to the limitations described in the form of key-value pairs.")
		dms_describeRecommendationLimitationsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeRecommendationLimitationsCmd.Flags().String("next-token", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
	})
	dmsCmd.AddCommand(dms_describeRecommendationLimitationsCmd)
}
