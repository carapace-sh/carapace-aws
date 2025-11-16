package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeImagesCmd = &cobra.Command{
	Use:   "describe-images",
	Short: "Retrieves a list that describes one or more specified images, if the image names or image ARNs are provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeImagesCmd).Standalone()

	appstream_describeImagesCmd.Flags().String("arns", "", "The ARNs of the public, private, and shared images to describe.")
	appstream_describeImagesCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeImagesCmd.Flags().String("names", "", "The names of the public or private images to describe.")
	appstream_describeImagesCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstream_describeImagesCmd.Flags().String("type", "", "The type of image (public, private, or shared) to describe.")
	appstreamCmd.AddCommand(appstream_describeImagesCmd)
}
