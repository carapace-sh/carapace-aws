package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Deletes a model-specific endpoint for a previously-trained custom model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_deleteEndpointCmd).Standalone()

	comprehend_deleteEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Number (ARN) of the endpoint being deleted.")
	comprehend_deleteEndpointCmd.MarkFlagRequired("endpoint-arn")
	comprehendCmd.AddCommand(comprehend_deleteEndpointCmd)
}
