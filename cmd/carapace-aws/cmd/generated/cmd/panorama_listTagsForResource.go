package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listTagsForResourceCmd).Standalone()

		panorama_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource's ARN.")
		panorama_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	panoramaCmd.AddCommand(panorama_listTagsForResourceCmd)
}
