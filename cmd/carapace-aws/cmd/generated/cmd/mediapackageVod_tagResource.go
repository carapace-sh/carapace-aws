package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_tagResourceCmd).Standalone()

	mediapackageVod_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource.")
	mediapackageVod_tagResourceCmd.Flags().String("tags", "", "A collection of tags associated with a resource")
	mediapackageVod_tagResourceCmd.MarkFlagRequired("resource-arn")
	mediapackageVod_tagResourceCmd.MarkFlagRequired("tags")
	mediapackageVodCmd.AddCommand(mediapackageVod_tagResourceCmd)
}
