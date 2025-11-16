package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listContextsCmd = &cobra.Command{
	Use:   "list-contexts",
	Short: "Lists the contexts in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listContextsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listContextsCmd).Standalone()

		sagemaker_listContextsCmd.Flags().String("context-type", "", "A filter that returns only contexts of the specified type.")
		sagemaker_listContextsCmd.Flags().String("created-after", "", "A filter that returns only contexts created on or after the specified time.")
		sagemaker_listContextsCmd.Flags().String("created-before", "", "A filter that returns only contexts created on or before the specified time.")
		sagemaker_listContextsCmd.Flags().String("max-results", "", "The maximum number of contexts to return in the response.")
		sagemaker_listContextsCmd.Flags().String("next-token", "", "If the previous call to `ListContexts` didn't return the full set of contexts, the call returns a token for getting the next set of contexts.")
		sagemaker_listContextsCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listContextsCmd.Flags().String("sort-order", "", "The sort order.")
		sagemaker_listContextsCmd.Flags().String("source-uri", "", "A filter that returns only contexts with the specified source URI.")
	})
	sagemakerCmd.AddCommand(sagemaker_listContextsCmd)
}
