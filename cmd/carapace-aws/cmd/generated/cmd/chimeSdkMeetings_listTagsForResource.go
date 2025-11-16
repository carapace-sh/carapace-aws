package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMeetings_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags available for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMeetings_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMeetings_listTagsForResourceCmd).Standalone()

		chimeSdkMeetings_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		chimeSdkMeetings_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	chimeSdkMeetingsCmd.AddCommand(chimeSdkMeetings_listTagsForResourceCmd)
}
