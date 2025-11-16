package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessaging_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies the specified tags to the specified Amazon Chime SDK messaging resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessaging_tagResourceCmd).Standalone()

	chimeSdkMessaging_tagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
	chimeSdkMessaging_tagResourceCmd.Flags().String("tags", "", "The tag key-value pairs.")
	chimeSdkMessaging_tagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMessaging_tagResourceCmd.MarkFlagRequired("tags")
	chimeSdkMessagingCmd.AddCommand(chimeSdkMessaging_tagResourceCmd)
}
