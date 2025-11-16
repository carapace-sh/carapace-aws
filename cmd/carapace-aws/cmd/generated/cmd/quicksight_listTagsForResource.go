package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listTagsForResourceCmd).Standalone()

		quicksight_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want a list of tags for.")
		quicksight_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	quicksightCmd.AddCommand(quicksight_listTagsForResourceCmd)
}
