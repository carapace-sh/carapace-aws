package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a bot, bot alias, or bot channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_untagResourceCmd).Standalone()

		lexv2Models_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove the tags from.")
		lexv2Models_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the resource.")
		lexv2Models_untagResourceCmd.MarkFlagRequired("resource-arn")
		lexv2Models_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_untagResourceCmd)
}
