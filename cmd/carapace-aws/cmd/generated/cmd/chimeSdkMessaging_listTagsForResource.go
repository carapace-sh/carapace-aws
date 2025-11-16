package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags applied to an Amazon Chime SDK messaging resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_listTagsForResourceCmd).Standalone()

	chimeSdkMessaging_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	chimeSdkMessaging_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_listTagsForResourceCmd)
}
