package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags associated with a notification rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_listTagsForResourceCmd).Standalone()

	codestarNotifications_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for the notification rule.")
	codestarNotifications_listTagsForResourceCmd.MarkFlagRequired("arn")
	codestarNotificationsCmd.AddCommand(codestarNotifications_listTagsForResourceCmd)
}
