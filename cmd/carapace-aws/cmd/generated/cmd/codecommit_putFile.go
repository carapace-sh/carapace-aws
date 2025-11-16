package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_putFileCmd = &cobra.Command{
	Use:   "put-file",
	Short: "Adds or updates a file in a branch in an CodeCommit repository, and generates a commit for the addition in the specified branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_putFileCmd).Standalone()

	codecommit_putFileCmd.Flags().String("branch-name", "", "The name of the branch where you want to add or update the file.")
	codecommit_putFileCmd.Flags().String("commit-message", "", "A message about why this file was added or updated.")
	codecommit_putFileCmd.Flags().String("email", "", "An email address for the person adding or updating the file.")
	codecommit_putFileCmd.Flags().String("file-content", "", "The content of the file, in binary object format.")
	codecommit_putFileCmd.Flags().String("file-mode", "", "The file mode permissions of the blob.")
	codecommit_putFileCmd.Flags().String("file-path", "", "The name of the file you want to add or update, including the relative path to the file in the repository.")
	codecommit_putFileCmd.Flags().String("name", "", "The name of the person adding or updating the file.")
	codecommit_putFileCmd.Flags().String("parent-commit-id", "", "The full commit ID of the head commit in the branch where you want to add or update the file.")
	codecommit_putFileCmd.Flags().String("repository-name", "", "The name of the repository where you want to add or update the file.")
	codecommit_putFileCmd.MarkFlagRequired("branch-name")
	codecommit_putFileCmd.MarkFlagRequired("file-content")
	codecommit_putFileCmd.MarkFlagRequired("file-path")
	codecommit_putFileCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_putFileCmd)
}
