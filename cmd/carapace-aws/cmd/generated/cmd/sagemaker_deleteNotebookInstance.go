package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteNotebookInstanceCmd = &cobra.Command{
	Use:   "delete-notebook-instance",
	Short: "Deletes an SageMaker AI notebook instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteNotebookInstanceCmd).Standalone()

	sagemaker_deleteNotebookInstanceCmd.Flags().String("notebook-instance-name", "", "The name of the SageMaker AI notebook instance to delete.")
	sagemaker_deleteNotebookInstanceCmd.MarkFlagRequired("notebook-instance-name")
	sagemakerCmd.AddCommand(sagemaker_deleteNotebookInstanceCmd)
}
