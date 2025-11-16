package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_untagResourceCmd).Standalone()

		quicksight_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
		quicksight_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the key-value pairs for the resource tag or tags assigned to the resource.")
		quicksight_untagResourceCmd.MarkFlagRequired("resource-arn")
		quicksight_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	quicksightCmd.AddCommand(quicksight_untagResourceCmd)
}
