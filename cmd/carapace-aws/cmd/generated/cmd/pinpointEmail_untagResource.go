package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove one or more tags (keys and values) from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_untagResourceCmd).Standalone()

	pinpointEmail_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove one or more tags from.")
	pinpointEmail_untagResourceCmd.Flags().String("tag-keys", "", "The tags (tag keys) that you want to remove from the resource.")
	pinpointEmail_untagResourceCmd.MarkFlagRequired("resource-arn")
	pinpointEmail_untagResourceCmd.MarkFlagRequired("tag-keys")
	pinpointEmailCmd.AddCommand(pinpointEmail_untagResourceCmd)
}
