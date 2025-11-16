package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getCommitCmd = &cobra.Command{
	Use:   "get-commit",
	Short: "Returns information about a commit, including commit message and committer information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getCommitCmd).Standalone()

	codecommit_getCommitCmd.Flags().String("commit-id", "", "The commit ID.")
	codecommit_getCommitCmd.Flags().String("repository-name", "", "The name of the repository to which the commit was made.")
	codecommit_getCommitCmd.MarkFlagRequired("commit-id")
	codecommit_getCommitCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getCommitCmd)
}
