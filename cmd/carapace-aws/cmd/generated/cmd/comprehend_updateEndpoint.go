package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_updateEndpointCmd = &cobra.Command{
	Use:   "update-endpoint",
	Short: "Updates information about the specified endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_updateEndpointCmd).Standalone()

	comprehend_updateEndpointCmd.Flags().String("desired-data-access-role-arn", "", "Data access role ARN to use in case the new model is encrypted with a customer CMK.")
	comprehend_updateEndpointCmd.Flags().String("desired-inference-units", "", "The desired number of inference units to be used by the model using this endpoint.")
	comprehend_updateEndpointCmd.Flags().String("desired-model-arn", "", "The ARN of the new model to use when updating an existing endpoint.")
	comprehend_updateEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Number (ARN) of the endpoint being updated.")
	comprehend_updateEndpointCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel")
	comprehend_updateEndpointCmd.MarkFlagRequired("endpoint-arn")
	comprehendCmd.AddCommand(comprehend_updateEndpointCmd)
}
