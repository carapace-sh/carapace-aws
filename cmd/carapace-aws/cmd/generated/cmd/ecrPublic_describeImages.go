package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_describeImagesCmd = &cobra.Command{
	Use:   "describe-images",
	Short: "Returns metadata that's related to the images in a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_describeImagesCmd).Standalone()

	ecrPublic_describeImagesCmd.Flags().String("image-ids", "", "The list of image IDs for the requested repository.")
	ecrPublic_describeImagesCmd.Flags().String("max-results", "", "The maximum number of repository results that's returned by `DescribeImages` in paginated output.")
	ecrPublic_describeImagesCmd.Flags().String("next-token", "", "The `nextToken` value that's returned from a previous paginated `DescribeImages` request where `maxResults` was used and the results exceeded the value of that parameter.")
	ecrPublic_describeImagesCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the public registry that contains the repository where images are described.")
	ecrPublic_describeImagesCmd.Flags().String("repository-name", "", "The repository that contains the images to describe.")
	ecrPublic_describeImagesCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_describeImagesCmd)
}
