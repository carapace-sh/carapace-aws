package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelCmd = &cobra.Command{
	Use:   "delete-model",
	Short: "Deletes a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteModelCmd).Standalone()

		sagemaker_deleteModelCmd.Flags().String("model-name", "", "The name of the model to delete.")
		sagemaker_deleteModelCmd.MarkFlagRequired("model-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteModelCmd)
}
