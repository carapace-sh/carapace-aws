package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Gets a list of the projects in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listProjectsCmd).Standalone()

	sagemaker_listProjectsCmd.Flags().String("creation-time-after", "", "A filter that returns the projects that were created after a specified time.")
	sagemaker_listProjectsCmd.Flags().String("creation-time-before", "", "A filter that returns the projects that were created before a specified time.")
	sagemaker_listProjectsCmd.Flags().String("max-results", "", "The maximum number of projects to return in the response.")
	sagemaker_listProjectsCmd.Flags().String("name-contains", "", "A filter that returns the projects whose name contains a specified string.")
	sagemaker_listProjectsCmd.Flags().String("next-token", "", "If the result of the previous `ListProjects` request was truncated, the response includes a `NextToken`.")
	sagemaker_listProjectsCmd.Flags().String("sort-by", "", "The field by which to sort results.")
	sagemaker_listProjectsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemakerCmd.AddCommand(sagemaker_listProjectsCmd)
}
