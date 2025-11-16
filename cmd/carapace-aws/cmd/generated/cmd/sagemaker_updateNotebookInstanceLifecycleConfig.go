package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateNotebookInstanceLifecycleConfigCmd = &cobra.Command{
	Use:   "update-notebook-instance-lifecycle-config",
	Short: "Updates a notebook instance lifecycle configuration created with the [CreateNotebookInstanceLifecycleConfig](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateNotebookInstanceLifecycleConfig.html) API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateNotebookInstanceLifecycleConfigCmd).Standalone()

	sagemaker_updateNotebookInstanceLifecycleConfigCmd.Flags().String("notebook-instance-lifecycle-config-name", "", "The name of the lifecycle configuration.")
	sagemaker_updateNotebookInstanceLifecycleConfigCmd.Flags().String("on-create", "", "The shell script that runs only once, when you create a notebook instance.")
	sagemaker_updateNotebookInstanceLifecycleConfigCmd.Flags().String("on-start", "", "The shell script that runs every time you start a notebook instance, including when you create the notebook instance.")
	sagemaker_updateNotebookInstanceLifecycleConfigCmd.MarkFlagRequired("notebook-instance-lifecycle-config-name")
	sagemakerCmd.AddCommand(sagemaker_updateNotebookInstanceLifecycleConfigCmd)
}
