package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listEndpointConfigsCmd = &cobra.Command{
	Use:   "list-endpoint-configs",
	Short: "Lists endpoint configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listEndpointConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listEndpointConfigsCmd).Standalone()

		sagemaker_listEndpointConfigsCmd.Flags().String("creation-time-after", "", "A filter that returns only endpoint configurations with a creation time greater than or equal to the specified time (timestamp).")
		sagemaker_listEndpointConfigsCmd.Flags().String("creation-time-before", "", "A filter that returns only endpoint configurations created before the specified time (timestamp).")
		sagemaker_listEndpointConfigsCmd.Flags().String("max-results", "", "The maximum number of training jobs to return in the response.")
		sagemaker_listEndpointConfigsCmd.Flags().String("name-contains", "", "A string in the endpoint configuration name.")
		sagemaker_listEndpointConfigsCmd.Flags().String("next-token", "", "If the result of the previous `ListEndpointConfig` request was truncated, the response includes a `NextToken`.")
		sagemaker_listEndpointConfigsCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listEndpointConfigsCmd.Flags().String("sort-order", "", "The sort order for results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listEndpointConfigsCmd)
}
