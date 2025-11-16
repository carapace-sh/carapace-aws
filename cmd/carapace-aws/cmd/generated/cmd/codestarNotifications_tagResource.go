package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a set of provided tags with a notification rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_tagResourceCmd).Standalone()

	codestarNotifications_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the notification rule to tag.")
	codestarNotifications_tagResourceCmd.Flags().String("tags", "", "The list of tags to associate with the resource.")
	codestarNotifications_tagResourceCmd.MarkFlagRequired("arn")
	codestarNotifications_tagResourceCmd.MarkFlagRequired("tags")
	codestarNotificationsCmd.AddCommand(codestarNotifications_tagResourceCmd)
}
