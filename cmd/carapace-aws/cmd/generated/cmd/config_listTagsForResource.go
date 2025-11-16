package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags for Config resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_listTagsForResourceCmd).Standalone()

		config_listTagsForResourceCmd.Flags().String("limit", "", "The maximum number of tags returned on each page.")
		config_listTagsForResourceCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
		config_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	configCmd.AddCommand(config_listTagsForResourceCmd)
}
