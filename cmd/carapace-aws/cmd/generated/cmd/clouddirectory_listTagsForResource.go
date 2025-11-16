package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listTagsForResourceCmd).Standalone()

	clouddirectory_listTagsForResourceCmd.Flags().String("max-results", "", "The `MaxResults` parameter sets the maximum number of results returned in a single page.")
	clouddirectory_listTagsForResourceCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	clouddirectory_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_listTagsForResourceCmd)
}
