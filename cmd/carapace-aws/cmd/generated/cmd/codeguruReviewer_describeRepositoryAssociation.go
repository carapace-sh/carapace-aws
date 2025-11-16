package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewer_describeRepositoryAssociationCmd = &cobra.Command{
	Use:   "describe-repository-association",
	Short: "Returns a [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object that contains information about the requested repository association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewer_describeRepositoryAssociationCmd).Standalone()

	codeguruReviewer_describeRepositoryAssociationCmd.Flags().String("association-arn", "", "The Amazon Resource Name (ARN) of the [RepositoryAssociation](https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html) object.")
	codeguruReviewer_describeRepositoryAssociationCmd.MarkFlagRequired("association-arn")
	codeguruReviewerCmd.AddCommand(codeguruReviewer_describeRepositoryAssociationCmd)
}
