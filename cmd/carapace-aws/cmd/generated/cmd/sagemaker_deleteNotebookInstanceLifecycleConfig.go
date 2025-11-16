package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteNotebookInstanceLifecycleConfigCmd = &cobra.Command{
	Use:   "delete-notebook-instance-lifecycle-config",
	Short: "Deletes a notebook instance lifecycle configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteNotebookInstanceLifecycleConfigCmd).Standalone()

	sagemaker_deleteNotebookInstanceLifecycleConfigCmd.Flags().String("notebook-instance-lifecycle-config-name", "", "The name of the lifecycle configuration to delete.")
	sagemaker_deleteNotebookInstanceLifecycleConfigCmd.MarkFlagRequired("notebook-instance-lifecycle-config-name")
	sagemakerCmd.AddCommand(sagemaker_deleteNotebookInstanceLifecycleConfigCmd)
}
