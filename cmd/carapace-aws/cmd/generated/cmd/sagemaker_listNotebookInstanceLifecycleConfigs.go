package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listNotebookInstanceLifecycleConfigsCmd = &cobra.Command{
	Use:   "list-notebook-instance-lifecycle-configs",
	Short: "Lists notebook instance lifestyle configurations created with the [CreateNotebookInstanceLifecycleConfig](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateNotebookInstanceLifecycleConfig.html) API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listNotebookInstanceLifecycleConfigsCmd).Standalone()

	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("creation-time-after", "", "A filter that returns only lifecycle configurations that were created after the specified time (timestamp).")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("creation-time-before", "", "A filter that returns only lifecycle configurations that were created before the specified time (timestamp).")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("last-modified-time-after", "", "A filter that returns only lifecycle configurations that were modified after the specified time (timestamp).")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("last-modified-time-before", "", "A filter that returns only lifecycle configurations that were modified before the specified time (timestamp).")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("max-results", "", "The maximum number of lifecycle configurations to return in the response.")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("name-contains", "", "A string in the lifecycle configuration name.")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("next-token", "", "If the result of a `ListNotebookInstanceLifecycleConfigs` request was truncated, the response includes a `NextToken`.")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("sort-by", "", "Sorts the list of results.")
	sagemaker_listNotebookInstanceLifecycleConfigsCmd.Flags().String("sort-order", "", "The sort order for results.")
	sagemakerCmd.AddCommand(sagemaker_listNotebookInstanceLifecycleConfigsCmd)
}
