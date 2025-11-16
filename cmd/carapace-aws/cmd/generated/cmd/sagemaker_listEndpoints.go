package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listEndpointsCmd = &cobra.Command{
	Use:   "list-endpoints",
	Short: "Lists endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listEndpointsCmd).Standalone()

	sagemaker_listEndpointsCmd.Flags().String("creation-time-after", "", "A filter that returns only endpoints with a creation time greater than or equal to the specified time (timestamp).")
	sagemaker_listEndpointsCmd.Flags().String("creation-time-before", "", "A filter that returns only endpoints that were created before the specified time (timestamp).")
	sagemaker_listEndpointsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only endpoints that were modified after the specified timestamp.")
	sagemaker_listEndpointsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only endpoints that were modified before the specified timestamp.")
	sagemaker_listEndpointsCmd.Flags().String("max-results", "", "The maximum number of endpoints to return in the response.")
	sagemaker_listEndpointsCmd.Flags().String("name-contains", "", "A string in endpoint names.")
	sagemaker_listEndpointsCmd.Flags().String("next-token", "", "If the result of a `ListEndpoints` request was truncated, the response includes a `NextToken`.")
	sagemaker_listEndpointsCmd.Flags().String("sort-by", "", "Sorts the list of results.")
	sagemaker_listEndpointsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemaker_listEndpointsCmd.Flags().String("status-equals", "", "A filter that returns only endpoints with the specified status.")
	sagemakerCmd.AddCommand(sagemaker_listEndpointsCmd)
}
