package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_listRepositoryAssociationsCmd = &cobra.Command{
	Use:   "list-repository-associations",
	Short: "Returns a list of [RepositoryAssociationSummary](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html) objects that contain summary information about a repository association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_listRepositoryAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_listRepositoryAssociationsCmd).Standalone()

		codeguruReviewer_listRepositoryAssociationsCmd.Flags().String("max-results", "", "The maximum number of repository association results returned by `ListRepositoryAssociations` in paginated output.")
		codeguruReviewer_listRepositoryAssociationsCmd.Flags().String("names", "", "List of repository names to use as a filter.")
		codeguruReviewer_listRepositoryAssociationsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListRepositoryAssociations` request where `maxResults` was used and the results exceeded the value of that parameter.")
		codeguruReviewer_listRepositoryAssociationsCmd.Flags().String("owners", "", "List of owners to use as a filter.")
		codeguruReviewer_listRepositoryAssociationsCmd.Flags().String("provider-types", "", "List of provider types to use as a filter.")
		codeguruReviewer_listRepositoryAssociationsCmd.Flags().String("states", "", "List of repository association states to use as a filter.")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_listRepositoryAssociationsCmd)
}
