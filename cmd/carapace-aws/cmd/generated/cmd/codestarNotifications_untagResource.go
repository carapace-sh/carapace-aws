package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association between one or more provided tags and a notification rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarNotifications_untagResourceCmd).Standalone()

		codestarNotifications_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule from which to remove the tags.")
		codestarNotifications_untagResourceCmd.Flags().String("tag-keys", "", "The key names of the tags to remove.")
		codestarNotifications_untagResourceCmd.MarkFlagRequired("arn")
		codestarNotifications_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codestarNotificationsCmd.AddCommand(codestarNotifications_untagResourceCmd)
}
