package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags the resource with a tag key and value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_tagResourceCmd).Standalone()

	notifications_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) to use to tag a resource.")
	notifications_tagResourceCmd.Flags().String("tags", "", "A map of tags assigned to a resource.")
	notifications_tagResourceCmd.MarkFlagRequired("arn")
	notifications_tagResourceCmd.MarkFlagRequired("tags")
	notificationsCmd.AddCommand(notifications_tagResourceCmd)
}
