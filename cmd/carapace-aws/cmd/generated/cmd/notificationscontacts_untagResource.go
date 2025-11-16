package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Detaches a key-value pair from a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontacts_untagResourceCmd).Standalone()

		notificationscontacts_untagResourceCmd.Flags().String("arn", "", "The value of the resource that will have the tag removed.")
		notificationscontacts_untagResourceCmd.Flags().String("tag-keys", "", "Specifies a list of tag keys that you want to remove from the specified resources.")
		notificationscontacts_untagResourceCmd.MarkFlagRequired("arn")
		notificationscontacts_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	notificationscontactsCmd.AddCommand(notificationscontacts_untagResourceCmd)
}
