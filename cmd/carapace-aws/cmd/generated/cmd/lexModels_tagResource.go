package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_tagResourceCmd).Standalone()

	lexModels_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot, bot alias, or bot channel to tag.")
	lexModels_tagResourceCmd.Flags().String("tags", "", "A list of tag keys to add to the resource.")
	lexModels_tagResourceCmd.MarkFlagRequired("resource-arn")
	lexModels_tagResourceCmd.MarkFlagRequired("tags")
	lexModelsCmd.AddCommand(lexModels_tagResourceCmd)
}
