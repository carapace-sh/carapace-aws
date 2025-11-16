package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags that have been added to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listTagsForResourceCmd).Standalone()

		cleanrooms_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource you want to list tags on.")
		cleanrooms_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listTagsForResourceCmd)
}
