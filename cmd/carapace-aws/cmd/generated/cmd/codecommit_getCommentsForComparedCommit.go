package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getCommentsForComparedCommitCmd = &cobra.Command{
	Use:   "get-comments-for-compared-commit",
	Short: "Returns information about comments made on the comparison between two commits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getCommentsForComparedCommitCmd).Standalone()

	codecommit_getCommentsForComparedCommitCmd.Flags().String("after-commit-id", "", "To establish the directionality of the comparison, the full commit ID of the after commit.")
	codecommit_getCommentsForComparedCommitCmd.Flags().String("before-commit-id", "", "To establish the directionality of the comparison, the full commit ID of the before commit.")
	codecommit_getCommentsForComparedCommitCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codecommit_getCommentsForComparedCommitCmd.Flags().String("next-token", "", "An enumeration token that when provided in a request, returns the next batch of the results.")
	codecommit_getCommentsForComparedCommitCmd.Flags().String("repository-name", "", "The name of the repository where you want to compare commits.")
	codecommit_getCommentsForComparedCommitCmd.MarkFlagRequired("after-commit-id")
	codecommit_getCommentsForComparedCommitCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getCommentsForComparedCommitCmd)
}
