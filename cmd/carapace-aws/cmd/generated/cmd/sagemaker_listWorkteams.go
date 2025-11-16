package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listWorkteamsCmd = &cobra.Command{
	Use:   "list-workteams",
	Short: "Gets a list of private work teams that you have defined in a region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listWorkteamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listWorkteamsCmd).Standalone()

		sagemaker_listWorkteamsCmd.Flags().String("max-results", "", "The maximum number of work teams to return in each page of the response.")
		sagemaker_listWorkteamsCmd.Flags().String("name-contains", "", "A string in the work team's name.")
		sagemaker_listWorkteamsCmd.Flags().String("next-token", "", "If the result of the previous `ListWorkteams` request was truncated, the response includes a `NextToken`.")
		sagemaker_listWorkteamsCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listWorkteamsCmd.Flags().String("sort-order", "", "The sort order for results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listWorkteamsCmd)
}
