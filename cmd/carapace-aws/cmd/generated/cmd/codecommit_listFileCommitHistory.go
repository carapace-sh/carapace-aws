package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listFileCommitHistoryCmd = &cobra.Command{
	Use:   "list-file-commit-history",
	Short: "Retrieves a list of commits and changes to a specified file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listFileCommitHistoryCmd).Standalone()

	codecommit_listFileCommitHistoryCmd.Flags().String("commit-specifier", "", "The fully quaified reference that identifies the commit that contains the file.")
	codecommit_listFileCommitHistoryCmd.Flags().String("file-path", "", "The full path of the file whose history you want to retrieve, including the name of the file.")
	codecommit_listFileCommitHistoryCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codecommit_listFileCommitHistoryCmd.Flags().String("next-token", "", "An enumeration token that allows the operation to batch the results.")
	codecommit_listFileCommitHistoryCmd.Flags().String("repository-name", "", "The name of the repository that contains the file.")
	codecommit_listFileCommitHistoryCmd.MarkFlagRequired("file-path")
	codecommit_listFileCommitHistoryCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_listFileCommitHistoryCmd)
}
