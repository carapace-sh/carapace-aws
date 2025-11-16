package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_listImagesCmd = &cobra.Command{
	Use:   "list-images",
	Short: "Lists all the image IDs for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_listImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_listImagesCmd).Standalone()

		ecr_listImagesCmd.Flags().String("filter", "", "The filter key and value with which to filter your `ListImages` results.")
		ecr_listImagesCmd.Flags().String("max-results", "", "The maximum number of image results returned by `ListImages` in paginated output.")
		ecr_listImagesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListImages` request where `maxResults` was used and the results exceeded the value of that parameter.")
		ecr_listImagesCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to list images.")
		ecr_listImagesCmd.Flags().String("repository-name", "", "The repository with image IDs to be listed.")
		ecr_listImagesCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_listImagesCmd)
}
