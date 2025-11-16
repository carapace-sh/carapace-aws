package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_describeBucketsCmd = &cobra.Command{
	Use:   "describe-buckets",
	Short: "Retrieves (queries) statistical data and other information about one or more S3 buckets that Amazon Macie monitors and analyzes for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_describeBucketsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_describeBucketsCmd).Standalone()

		macie2_describeBucketsCmd.Flags().String("criteria", "", "The criteria to use to filter the query results.")
		macie2_describeBucketsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of the response.")
		macie2_describeBucketsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
		macie2_describeBucketsCmd.Flags().String("sort-criteria", "", "The criteria to use to sort the query results.")
	})
	macie2Cmd.AddCommand(macie2_describeBucketsCmd)
}
