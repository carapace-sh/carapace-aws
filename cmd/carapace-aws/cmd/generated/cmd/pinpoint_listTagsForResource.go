package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves all the tags (keys and values) that are associated with an application, campaign, message template, or segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_listTagsForResourceCmd).Standalone()

	pinpoint_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	pinpoint_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	pinpointCmd.AddCommand(pinpoint_listTagsForResourceCmd)
}
