package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelCardCmd = &cobra.Command{
	Use:   "delete-model-card",
	Short: "Deletes an Amazon SageMaker Model Card.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelCardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteModelCardCmd).Standalone()

		sagemaker_deleteModelCardCmd.Flags().String("model-card-name", "", "The name of the model card to delete.")
		sagemaker_deleteModelCardCmd.MarkFlagRequired("model-card-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteModelCardCmd)
}
