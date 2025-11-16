package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeCodeRepositoryCmd = &cobra.Command{
	Use:   "describe-code-repository",
	Short: "Gets details about the specified Git repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeCodeRepositoryCmd).Standalone()

	sagemaker_describeCodeRepositoryCmd.Flags().String("code-repository-name", "", "The name of the Git repository to describe.")
	sagemaker_describeCodeRepositoryCmd.MarkFlagRequired("code-repository-name")
	sagemakerCmd.AddCommand(sagemaker_describeCodeRepositoryCmd)
}
