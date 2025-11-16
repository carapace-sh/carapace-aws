package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createCodeRepositoryCmd = &cobra.Command{
	Use:   "create-code-repository",
	Short: "Creates a Git repository as a resource in your SageMaker AI account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createCodeRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createCodeRepositoryCmd).Standalone()

		sagemaker_createCodeRepositoryCmd.Flags().String("code-repository-name", "", "The name of the Git repository.")
		sagemaker_createCodeRepositoryCmd.Flags().String("git-config", "", "Specifies details about the repository, including the URL where the repository is located, the default branch, and credentials to use to access the repository.")
		sagemaker_createCodeRepositoryCmd.Flags().String("tags", "", "An array of key-value pairs.")
		sagemaker_createCodeRepositoryCmd.MarkFlagRequired("code-repository-name")
		sagemaker_createCodeRepositoryCmd.MarkFlagRequired("git-config")
	})
	sagemakerCmd.AddCommand(sagemaker_createCodeRepositoryCmd)
}
