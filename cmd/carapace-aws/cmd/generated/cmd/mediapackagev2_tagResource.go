package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one of more tags (key-value pairs) to the specified MediaPackage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_tagResourceCmd).Standalone()

		mediapackagev2_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the MediaPackage resource that you're adding tags to.")
		mediapackagev2_tagResourceCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.")
		mediapackagev2_tagResourceCmd.MarkFlagRequired("resource-arn")
		mediapackagev2_tagResourceCmd.MarkFlagRequired("tags")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_tagResourceCmd)
}
