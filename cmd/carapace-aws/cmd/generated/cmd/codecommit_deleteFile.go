package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_deleteFileCmd = &cobra.Command{
	Use:   "delete-file",
	Short: "Deletes a specified file from a specified branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_deleteFileCmd).Standalone()

	codecommit_deleteFileCmd.Flags().String("branch-name", "", "The name of the branch where the commit that deletes the file is made.")
	codecommit_deleteFileCmd.Flags().String("commit-message", "", "The commit message you want to include as part of deleting the file.")
	codecommit_deleteFileCmd.Flags().String("email", "", "The email address for the commit that deletes the file.")
	codecommit_deleteFileCmd.Flags().String("file-path", "", "The fully qualified path to the file that to be deleted, including the full name and extension of that file.")
	codecommit_deleteFileCmd.Flags().String("keep-empty-folders", "", "If a file is the only object in the folder or directory, specifies whether to delete the folder or directory that contains the file.")
	codecommit_deleteFileCmd.Flags().String("name", "", "The name of the author of the commit that deletes the file.")
	codecommit_deleteFileCmd.Flags().String("parent-commit-id", "", "The ID of the commit that is the tip of the branch where you want to create the commit that deletes the file.")
	codecommit_deleteFileCmd.Flags().String("repository-name", "", "The name of the repository that contains the file to delete.")
	codecommit_deleteFileCmd.MarkFlagRequired("branch-name")
	codecommit_deleteFileCmd.MarkFlagRequired("file-path")
	codecommit_deleteFileCmd.MarkFlagRequired("parent-commit-id")
	codecommit_deleteFileCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_deleteFileCmd)
}
