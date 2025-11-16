package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified Amazon Chime SDK messaging resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_untagResourceCmd).Standalone()

	chimeSdkMessaging_untagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
	chimeSdkMessaging_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
	chimeSdkMessaging_untagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMessaging_untagResourceCmd.MarkFlagRequired("tag-keys")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_untagResourceCmd)
}
