package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelExplainabilityJobDefinitionsCmd = &cobra.Command{
	Use:   "list-model-explainability-job-definitions",
	Short: "Lists model explainability job definitions that satisfy various filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelExplainabilityJobDefinitionsCmd).Standalone()

	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("creation-time-after", "", "A filter that returns only model explainability jobs created after a specified time.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("creation-time-before", "", "A filter that returns only model explainability jobs created before a specified time.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("endpoint-name", "", "Name of the endpoint to monitor for model explainability.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("max-results", "", "The maximum number of jobs to return in the response.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("name-contains", "", "Filter for model explainability jobs whose name contains a specified string.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("next-token", "", "The token returned if the response is truncated.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("sort-by", "", "Whether to sort results by the `Name` or `CreationTime` field.")
	sagemaker_listModelExplainabilityJobDefinitionsCmd.Flags().String("sort-order", "", "Whether to sort the results in `Ascending` or `Descending` order.")
	sagemakerCmd.AddCommand(sagemaker_listModelExplainabilityJobDefinitionsCmd)
}
