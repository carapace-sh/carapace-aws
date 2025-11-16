package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelCmd = &cobra.Command{
	Use:   "describe-model",
	Short: "Describes a model that you created using the `CreateModel` API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelCmd).Standalone()

	sagemaker_describeModelCmd.Flags().String("model-name", "", "The name of the model.")
	sagemaker_describeModelCmd.MarkFlagRequired("model-name")
	sagemakerCmd.AddCommand(sagemaker_describeModelCmd)
}
