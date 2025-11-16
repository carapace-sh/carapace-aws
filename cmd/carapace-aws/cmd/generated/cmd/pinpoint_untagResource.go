package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags (keys and values) from an application, campaign, message template, or segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_untagResourceCmd).Standalone()

	pinpoint_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	pinpoint_untagResourceCmd.Flags().String("tag-keys", "", "The key of the tag to remove from the resource.")
	pinpoint_untagResourceCmd.MarkFlagRequired("resource-arn")
	pinpoint_untagResourceCmd.MarkFlagRequired("tag-keys")
	pinpointCmd.AddCommand(pinpoint_untagResourceCmd)
}
