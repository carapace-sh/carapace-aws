package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Queries for available tag keys and tag values for a specified period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getTagsCmd).Standalone()

	ce_getTagsCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
	ce_getTagsCmd.Flags().String("filter", "", "")
	ce_getTagsCmd.Flags().String("max-results", "", "This field is only used when SortBy is provided in the request.")
	ce_getTagsCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getTagsCmd.Flags().String("search-string", "", "The value that you want to search for.")
	ce_getTagsCmd.Flags().String("sort-by", "", "The value that you want to sort the data by.")
	ce_getTagsCmd.Flags().String("tag-key", "", "The key of the tag that you want to return values for.")
	ce_getTagsCmd.Flags().String("time-period", "", "The start and end dates for retrieving the dimension values.")
	ce_getTagsCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getTagsCmd)
}
