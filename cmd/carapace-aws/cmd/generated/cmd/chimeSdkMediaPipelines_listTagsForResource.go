package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags available for a media pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_listTagsForResourceCmd).Standalone()

	chimeSdkMediaPipelines_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the media pipeline associated with any tags.")
	chimeSdkMediaPipelines_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_listTagsForResourceCmd)
}
