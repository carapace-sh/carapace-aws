package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags associated with an Athena resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_listTagsForResourceCmd).Standalone()

	athena_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request that lists the tags for the resource.")
	athena_listTagsForResourceCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no additional results for this request, where the request lists the tags for the resource with the specified ARN.")
	athena_listTagsForResourceCmd.Flags().String("resource-arn", "", "Lists the tags for the resource with the specified ARN.")
	athena_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	athenaCmd.AddCommand(athena_listTagsForResourceCmd)
}
