package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getCommentCmd = &cobra.Command{
	Use:   "get-comment",
	Short: "Returns the content of a comment made on a change, file, or commit in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getCommentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_getCommentCmd).Standalone()

		codecommit_getCommentCmd.Flags().String("comment-id", "", "The unique, system-generated ID of the comment.")
		codecommit_getCommentCmd.MarkFlagRequired("comment-id")
	})
	codecommitCmd.AddCommand(codecommit_getCommentCmd)
}
