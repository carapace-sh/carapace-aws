package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_untagResourceCmd).Standalone()

	chimeSdkMeetings_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you're removing tags from.")
	chimeSdkMeetings_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys being removed from the resources.")
	chimeSdkMeetings_untagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMeetings_untagResourceCmd.MarkFlagRequired("tag-keys")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_untagResourceCmd)
}
