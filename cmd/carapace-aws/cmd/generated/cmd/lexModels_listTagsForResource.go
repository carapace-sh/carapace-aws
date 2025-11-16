package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets a list of tags associated with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_listTagsForResourceCmd).Standalone()

	lexModels_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to get a list of tags for.")
	lexModels_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	lexModelsCmd.AddCommand(lexModels_listTagsForResourceCmd)
}
