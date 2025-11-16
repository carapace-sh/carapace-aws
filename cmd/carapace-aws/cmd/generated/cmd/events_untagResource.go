package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified EventBridge resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_untagResourceCmd).Standalone()

	events_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the EventBridge resource from which you are removing tags.")
	events_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	events_untagResourceCmd.MarkFlagRequired("resource-arn")
	events_untagResourceCmd.MarkFlagRequired("tag-keys")
	eventsCmd.AddCommand(events_untagResourceCmd)
}
