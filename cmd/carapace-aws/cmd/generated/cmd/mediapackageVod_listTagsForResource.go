package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags assigned to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_listTagsForResourceCmd).Standalone()

		mediapackageVod_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource.")
		mediapackageVod_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_listTagsForResourceCmd)
}
