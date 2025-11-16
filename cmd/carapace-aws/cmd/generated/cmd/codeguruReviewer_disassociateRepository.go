package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_disassociateRepositoryCmd = &cobra.Command{
	Use:   "disassociate-repository",
	Short: "Removes the association between Amazon CodeGuru Reviewer and a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_disassociateRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewer_disassociateRepositoryCmd).Standalone()

		codeguruReviewer_disassociateRepositoryCmd.Flags().String("association-arn", "", "The Amazon Resource Name (ARN) of the [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object.")
		codeguruReviewer_disassociateRepositoryCmd.MarkFlagRequired("association-arn")
	})
	codeguruReviewerCmd.AddCommand(codeguruReviewer_disassociateRepositoryCmd)
}
