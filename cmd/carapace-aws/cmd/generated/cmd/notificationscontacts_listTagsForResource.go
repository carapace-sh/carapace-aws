package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags associated with the Amazon Resource Name (ARN) that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_listTagsForResourceCmd).Standalone()

	notificationscontacts_listTagsForResourceCmd.Flags().String("arn", "", "The ARN you specified to list the tags of.")
	notificationscontacts_listTagsForResourceCmd.MarkFlagRequired("arn")
	notificationscontactsCmd.AddCommand(notificationscontacts_listTagsForResourceCmd)
}
