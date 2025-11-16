package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateCodeRepositoryCmd = &cobra.Command{
	Use:   "update-code-repository",
	Short: "Updates the specified Git repository with the specified values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateCodeRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateCodeRepositoryCmd).Standalone()

		sagemaker_updateCodeRepositoryCmd.Flags().String("code-repository-name", "", "The name of the Git repository to update.")
		sagemaker_updateCodeRepositoryCmd.Flags().String("git-config", "", "The configuration of the git repository, including the URL and the Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager secret that contains the credentials used to access the repository.")
		sagemaker_updateCodeRepositoryCmd.MarkFlagRequired("code-repository-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateCodeRepositoryCmd)
}
