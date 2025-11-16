package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_listTagsForResourceCmd).Standalone()

	mediapackagev2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch resource that you want to view tags for.")
	mediapackagev2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	mediapackagev2Cmd.AddCommand(mediapackagev2_listTagsForResourceCmd)
}
