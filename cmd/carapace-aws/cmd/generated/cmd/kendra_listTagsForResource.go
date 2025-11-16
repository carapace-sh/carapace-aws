package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets a list of tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_listTagsForResourceCmd).Standalone()

		kendra_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the index, FAQ, data source, or other resource to get a list of tags for.")
		kendra_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	kendraCmd.AddCommand(kendra_listTagsForResourceCmd)
}
