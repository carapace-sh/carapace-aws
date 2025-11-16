package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_searchResourcesCmd = &cobra.Command{
	Use:   "search-resources",
	Short: "Retrieves (queries) statistical data and other information about Amazon Web Services resources that Amazon Macie monitors and analyzes for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_searchResourcesCmd).Standalone()

	macie2_searchResourcesCmd.Flags().String("bucket-criteria", "", "The filter conditions that determine which S3 buckets to include or exclude from the query results.")
	macie2_searchResourcesCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of the response.")
	macie2_searchResourcesCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2_searchResourcesCmd.Flags().String("sort-criteria", "", "The criteria to use to sort the results.")
	macie2Cmd.AddCommand(macie2_searchResourcesCmd)
}
