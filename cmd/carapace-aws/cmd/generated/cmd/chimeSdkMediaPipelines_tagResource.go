package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "The ARN of the media pipeline that you want to tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_tagResourceCmd).Standalone()

	chimeSdkMediaPipelines_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the media pipeline associated with any tags.")
	chimeSdkMediaPipelines_tagResourceCmd.Flags().String("tags", "", "The tags associated with the specified media pipeline.")
	chimeSdkMediaPipelines_tagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMediaPipelines_tagResourceCmd.MarkFlagRequired("tags")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_tagResourceCmd)
}
