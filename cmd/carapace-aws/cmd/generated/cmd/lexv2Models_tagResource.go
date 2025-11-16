package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_tagResourceCmd).Standalone()

	lexv2Models_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot, bot alias, or bot channel to tag.")
	lexv2Models_tagResourceCmd.Flags().String("tags", "", "A list of tag keys to add to the resource.")
	lexv2Models_tagResourceCmd.MarkFlagRequired("resource-arn")
	lexv2Models_tagResourceCmd.MarkFlagRequired("tags")
	lexv2ModelsCmd.AddCommand(lexv2Models_tagResourceCmd)
}
