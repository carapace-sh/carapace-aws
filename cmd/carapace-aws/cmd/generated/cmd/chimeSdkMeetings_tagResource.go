package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "The resource that supports tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_tagResourceCmd).Standalone()

	chimeSdkMeetings_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	chimeSdkMeetings_tagResourceCmd.Flags().String("tags", "", "Lists the requested tags.")
	chimeSdkMeetings_tagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMeetings_tagResourceCmd.MarkFlagRequired("tags")
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_tagResourceCmd)
}
