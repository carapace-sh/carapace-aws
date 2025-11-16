package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describeImagesCmd = &cobra.Command{
	Use:   "describe-images",
	Short: "Returns metadata about the images in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describeImagesCmd).Standalone()

	ecr_describeImagesCmd.Flags().String("filter", "", "The filter key and value with which to filter your `DescribeImages` results.")
	ecr_describeImagesCmd.Flags().String("image-ids", "", "The list of image IDs for the requested repository.")
	ecr_describeImagesCmd.Flags().String("max-results", "", "The maximum number of repository results returned by `DescribeImages` in paginated output.")
	ecr_describeImagesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeImages` request where `maxResults` was used and the results exceeded the value of that parameter.")
	ecr_describeImagesCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to describe images.")
	ecr_describeImagesCmd.Flags().String("repository-name", "", "The repository that contains the images to describe.")
	ecr_describeImagesCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_describeImagesCmd)
}
