package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_associateRepositoryCmd = &cobra.Command{
	Use:   "associate-repository",
	Short: "Use to associate an Amazon Web Services CodeCommit repository or a repository managed by Amazon Web Services CodeStar Connections with Amazon CodeGuru Reviewer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_associateRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_associateRepositoryCmd).Standalone()

		codeguruReviewer_associateRepositoryCmd.Flags().String("client-request-token", "", "Amazon CodeGuru Reviewer uses this value to prevent the accidental creation of duplicate repository associations if there are failures and retries.")
		codeguruReviewer_associateRepositoryCmd.Flags().String("kmskey-details", "", "A `KMSKeyDetails` object that contains:")
		codeguruReviewer_associateRepositoryCmd.Flags().String("repository", "", "The repository to associate.")
		codeguruReviewer_associateRepositoryCmd.Flags().String("tags", "", "An array of key-value pairs used to tag an associated repository.")
		codeguruReviewer_associateRepositoryCmd.MarkFlagRequired("repository")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_associateRepositoryCmd)
}
