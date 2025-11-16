package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listCodeRepositoriesCmd = &cobra.Command{
	Use:   "list-code-repositories",
	Short: "Gets a list of the Git repositories in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listCodeRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listCodeRepositoriesCmd).Standalone()

		sagemaker_listCodeRepositoriesCmd.Flags().String("creation-time-after", "", "A filter that returns only Git repositories that were created after the specified time.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("creation-time-before", "", "A filter that returns only Git repositories that were created before the specified time.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("last-modified-time-after", "", "A filter that returns only Git repositories that were last modified after the specified time.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("last-modified-time-before", "", "A filter that returns only Git repositories that were last modified before the specified time.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("max-results", "", "The maximum number of Git repositories to return in the response.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("name-contains", "", "A string in the Git repositories name.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("next-token", "", "If the result of a `ListCodeRepositoriesOutput` request was truncated, the response includes a `NextToken`.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listCodeRepositoriesCmd.Flags().String("sort-order", "", "The sort order for results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listCodeRepositoriesCmd)
}
