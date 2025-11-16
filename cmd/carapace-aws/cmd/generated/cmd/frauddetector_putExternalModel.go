package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putExternalModelCmd = &cobra.Command{
	Use:   "put-external-model",
	Short: "Creates or updates an Amazon SageMaker model endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putExternalModelCmd).Standalone()

	frauddetector_putExternalModelCmd.Flags().String("input-configuration", "", "The model endpoint input configuration.")
	frauddetector_putExternalModelCmd.Flags().String("invoke-model-endpoint-role-arn", "", "The IAM role used to invoke the model endpoint.")
	frauddetector_putExternalModelCmd.Flags().String("model-endpoint", "", "The model endpoints name.")
	frauddetector_putExternalModelCmd.Flags().String("model-endpoint-status", "", "The model endpoint’s status in Amazon Fraud Detector.")
	frauddetector_putExternalModelCmd.Flags().String("model-source", "", "The source of the model.")
	frauddetector_putExternalModelCmd.Flags().String("output-configuration", "", "The model endpoint output configuration.")
	frauddetector_putExternalModelCmd.Flags().String("tags", "", "A collection of key and value pairs.")
	frauddetector_putExternalModelCmd.MarkFlagRequired("input-configuration")
	frauddetector_putExternalModelCmd.MarkFlagRequired("invoke-model-endpoint-role-arn")
	frauddetector_putExternalModelCmd.MarkFlagRequired("model-endpoint")
	frauddetector_putExternalModelCmd.MarkFlagRequired("model-endpoint-status")
	frauddetector_putExternalModelCmd.MarkFlagRequired("model-source")
	frauddetector_putExternalModelCmd.MarkFlagRequired("output-configuration")
	frauddetectorCmd.AddCommand(frauddetector_putExternalModelCmd)
}
