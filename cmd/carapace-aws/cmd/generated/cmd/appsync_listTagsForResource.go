package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listTagsForResourceCmd).Standalone()

		appsync_listTagsForResourceCmd.Flags().String("resource-arn", "", "The `GraphqlApi` Amazon Resource Name (ARN).")
		appsync_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	appsyncCmd.AddCommand(appsync_listTagsForResourceCmd)
}
