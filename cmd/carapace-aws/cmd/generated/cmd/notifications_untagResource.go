package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags a resource with a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_untagResourceCmd).Standalone()

		notifications_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) to use to untag a resource.")
		notifications_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to use to untag a resource.")
		notifications_untagResourceCmd.MarkFlagRequired("arn")
		notifications_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	notificationsCmd.AddCommand(notifications_untagResourceCmd)
}
