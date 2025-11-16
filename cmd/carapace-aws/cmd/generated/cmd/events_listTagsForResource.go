package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with an EventBridge resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listTagsForResourceCmd).Standalone()

	events_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the EventBridge resource for which you want to view tags.")
	events_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	eventsCmd.AddCommand(events_listTagsForResourceCmd)
}
