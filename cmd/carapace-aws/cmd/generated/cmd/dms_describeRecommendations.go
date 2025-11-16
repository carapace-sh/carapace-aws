package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeRecommendationsCmd = &cobra.Command{
	Use:   "describe-recommendations",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeRecommendationsCmd).Standalone()

	dms_describeRecommendationsCmd.Flags().String("filters", "", "Filters applied to the target engine recommendations described in the form of key-value pairs.")
	dms_describeRecommendationsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeRecommendationsCmd.Flags().String("next-token", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
	dmsCmd.AddCommand(dms_describeRecommendationsCmd)
}
