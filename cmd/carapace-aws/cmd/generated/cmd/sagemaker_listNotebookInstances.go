package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listNotebookInstancesCmd = &cobra.Command{
	Use:   "list-notebook-instances",
	Short: "Returns a list of the SageMaker AI notebook instances in the requester's account in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listNotebookInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listNotebookInstancesCmd).Standalone()

		sagemaker_listNotebookInstancesCmd.Flags().String("additional-code-repository-equals", "", "A filter that returns only notebook instances with associated with the specified git repository.")
		sagemaker_listNotebookInstancesCmd.Flags().String("creation-time-after", "", "A filter that returns only notebook instances that were created after the specified time (timestamp).")
		sagemaker_listNotebookInstancesCmd.Flags().String("creation-time-before", "", "A filter that returns only notebook instances that were created before the specified time (timestamp).")
		sagemaker_listNotebookInstancesCmd.Flags().String("default-code-repository-contains", "", "A string in the name or URL of a Git repository associated with this notebook instance.")
		sagemaker_listNotebookInstancesCmd.Flags().String("last-modified-time-after", "", "A filter that returns only notebook instances that were modified after the specified time (timestamp).")
		sagemaker_listNotebookInstancesCmd.Flags().String("last-modified-time-before", "", "A filter that returns only notebook instances that were modified before the specified time (timestamp).")
		sagemaker_listNotebookInstancesCmd.Flags().String("max-results", "", "The maximum number of notebook instances to return.")
		sagemaker_listNotebookInstancesCmd.Flags().String("name-contains", "", "A string in the notebook instances' name.")
		sagemaker_listNotebookInstancesCmd.Flags().String("next-token", "", "If the previous call to the `ListNotebookInstances` is truncated, the response includes a `NextToken`.")
		sagemaker_listNotebookInstancesCmd.Flags().String("notebook-instance-lifecycle-config-name-contains", "", "A string in the name of a notebook instances lifecycle configuration associated with this notebook instance.")
		sagemaker_listNotebookInstancesCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listNotebookInstancesCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listNotebookInstancesCmd.Flags().String("status-equals", "", "A filter that returns only notebook instances with the specified status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listNotebookInstancesCmd)
}
