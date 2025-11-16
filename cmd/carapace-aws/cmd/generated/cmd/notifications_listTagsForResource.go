package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notifications_listTagsForResourceCmd).Standalone()

		notifications_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) to use to list tags.")
		notifications_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	notificationsCmd.AddCommand(notifications_listTagsForResourceCmd)
}
