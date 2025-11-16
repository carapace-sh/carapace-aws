package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createCommitCmd = &cobra.Command{
	Use:   "create-commit",
	Short: "Creates a commit for a repository on the tip of a specified branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createCommitCmd).Standalone()

	codecommit_createCommitCmd.Flags().String("author-name", "", "The name of the author who created the commit.")
	codecommit_createCommitCmd.Flags().String("branch-name", "", "The name of the branch where you create the commit.")
	codecommit_createCommitCmd.Flags().String("commit-message", "", "The commit message you want to include in the commit.")
	codecommit_createCommitCmd.Flags().String("delete-files", "", "The files to delete in this commit.")
	codecommit_createCommitCmd.Flags().String("email", "", "The email address of the person who created the commit.")
	codecommit_createCommitCmd.Flags().String("keep-empty-folders", "", "If the commit contains deletions, whether to keep a folder or folder structure if the changes leave the folders empty.")
	codecommit_createCommitCmd.Flags().String("parent-commit-id", "", "The ID of the commit that is the parent of the commit you create.")
	codecommit_createCommitCmd.Flags().String("put-files", "", "The files to add or update in this commit.")
	codecommit_createCommitCmd.Flags().String("repository-name", "", "The name of the repository where you create the commit.")
	codecommit_createCommitCmd.Flags().String("set-file-modes", "", "The file modes to update for files in this commit.")
	codecommit_createCommitCmd.MarkFlagRequired("branch-name")
	codecommit_createCommitCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_createCommitCmd)
}
