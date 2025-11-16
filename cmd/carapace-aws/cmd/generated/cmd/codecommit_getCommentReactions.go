package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getCommentReactionsCmd = &cobra.Command{
	Use:   "get-comment-reactions",
	Short: "Returns information about reactions to a specified comment ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getCommentReactionsCmd).Standalone()

	codecommit_getCommentReactionsCmd.Flags().String("comment-id", "", "The ID of the comment for which you want to get reactions information.")
	codecommit_getCommentReactionsCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codecommit_getCommentReactionsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codecommit_getCommentReactionsCmd.Flags().String("reaction-user-arn", "", "Optional.")
	codecommit_getCommentReactionsCmd.MarkFlagRequired("comment-id")
	codecommitCmd.AddCommand(codecommit_getCommentReactionsCmd)
}
