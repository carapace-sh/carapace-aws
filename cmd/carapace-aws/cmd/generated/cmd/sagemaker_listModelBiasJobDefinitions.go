package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelBiasJobDefinitionsCmd = &cobra.Command{
	Use:   "list-model-bias-job-definitions",
	Short: "Lists model bias jobs definitions that satisfy various filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelBiasJobDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listModelBiasJobDefinitionsCmd).Standalone()

		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("creation-time-after", "", "A filter that returns only model bias jobs created after a specified time.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("creation-time-before", "", "A filter that returns only model bias jobs created before a specified time.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("endpoint-name", "", "Name of the endpoint to monitor for model bias.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("max-results", "", "The maximum number of model bias jobs to return in the response.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("name-contains", "", "Filter for model bias jobs whose name contains a specified string.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("next-token", "", "The token returned if the response is truncated.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("sort-by", "", "Whether to sort results by the `Name` or `CreationTime` field.")
		sagemaker_listModelBiasJobDefinitionsCmd.Flags().String("sort-order", "", "Whether to sort the results in `Ascending` or `Descending` order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listModelBiasJobDefinitionsCmd)
}
