package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches a key-value pair to a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_tagResourceCmd).Standalone()

	notificationscontacts_tagResourceCmd.Flags().String("arn", "", "The ARN of the configuration.")
	notificationscontacts_tagResourceCmd.Flags().String("tags", "", "A list of tags to apply to the configuration.")
	notificationscontacts_tagResourceCmd.MarkFlagRequired("arn")
	notificationscontacts_tagResourceCmd.MarkFlagRequired("tags")
	notificationscontactsCmd.AddCommand(notificationscontacts_tagResourceCmd)
}
