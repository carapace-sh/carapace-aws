package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_describeImageTagsCmd = &cobra.Command{
	Use:   "describe-image-tags",
	Short: "Returns the image tag details for a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_describeImageTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_describeImageTagsCmd).Standalone()

		ecrPublic_describeImageTagsCmd.Flags().String("max-results", "", "The maximum number of repository results that's returned by `DescribeImageTags` in paginated output.")
		ecrPublic_describeImageTagsCmd.Flags().String("next-token", "", "The `nextToken` value that's returned from a previous paginated `DescribeImageTags` request where `maxResults` was used and the results exceeded the value of that parameter.")
		ecrPublic_describeImageTagsCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the public registry that contains the repository where images are described.")
		ecrPublic_describeImageTagsCmd.Flags().String("repository-name", "", "The name of the repository that contains the image tag details to describe.")
		ecrPublic_describeImageTagsCmd.MarkFlagRequired("repository-name")
	})
	ecrPublicCmd.AddCommand(ecrPublic_describeImageTagsCmd)
}
