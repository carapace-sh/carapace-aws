package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove one or more tags (keys and values) from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_untagResourceCmd).Standalone()

	sesv2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove one or more tags from.")
	sesv2_untagResourceCmd.Flags().String("tag-keys", "", "The tags (tag keys) that you want to remove from the resource.")
	sesv2_untagResourceCmd.MarkFlagRequired("resource-arn")
	sesv2_untagResourceCmd.MarkFlagRequired("tag-keys")
	sesv2Cmd.AddCommand(sesv2_untagResourceCmd)
}
