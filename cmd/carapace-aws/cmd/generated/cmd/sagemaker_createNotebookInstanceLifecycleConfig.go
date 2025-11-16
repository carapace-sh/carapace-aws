package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createNotebookInstanceLifecycleConfigCmd = &cobra.Command{
	Use:   "create-notebook-instance-lifecycle-config",
	Short: "Creates a lifecycle configuration that you can associate with a notebook instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createNotebookInstanceLifecycleConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createNotebookInstanceLifecycleConfigCmd).Standalone()

		sagemaker_createNotebookInstanceLifecycleConfigCmd.Flags().String("notebook-instance-lifecycle-config-name", "", "The name of the lifecycle configuration.")
		sagemaker_createNotebookInstanceLifecycleConfigCmd.Flags().String("on-create", "", "A shell script that runs only once, when you create a notebook instance.")
		sagemaker_createNotebookInstanceLifecycleConfigCmd.Flags().String("on-start", "", "A shell script that runs every time you start a notebook instance, including when you create the notebook instance.")
		sagemaker_createNotebookInstanceLifecycleConfigCmd.Flags().String("tags", "", "An array of key-value pairs.")
		sagemaker_createNotebookInstanceLifecycleConfigCmd.MarkFlagRequired("notebook-instance-lifecycle-config-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createNotebookInstanceLifecycleConfigCmd)
}
