package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified EventBridge resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_tagResourceCmd).Standalone()

		events_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the EventBridge resource that you're adding tags to.")
		events_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
		events_tagResourceCmd.MarkFlagRequired("resource-arn")
		events_tagResourceCmd.MarkFlagRequired("tags")
	})
	eventsCmd.AddCommand(events_tagResourceCmd)
}
