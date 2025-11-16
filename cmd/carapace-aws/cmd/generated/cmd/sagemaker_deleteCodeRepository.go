package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteCodeRepositoryCmd = &cobra.Command{
	Use:   "delete-code-repository",
	Short: "Deletes the specified Git repository from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteCodeRepositoryCmd).Standalone()

	sagemaker_deleteCodeRepositoryCmd.Flags().String("code-repository-name", "", "The name of the Git repository to delete.")
	sagemaker_deleteCodeRepositoryCmd.MarkFlagRequired("code-repository-name")
	sagemakerCmd.AddCommand(sagemaker_deleteCodeRepositoryCmd)
}
