package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes any tags from a media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_untagResourceCmd).Standalone()

	chimeSdkMediaPipelines_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the pipeline that you want to untag.")
	chimeSdkMediaPipelines_untagResourceCmd.Flags().String("tag-keys", "", "The key/value pairs in the tag that you want to remove.")
	chimeSdkMediaPipelines_untagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMediaPipelines_untagResourceCmd.MarkFlagRequired("tag-keys")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_untagResourceCmd)
}
