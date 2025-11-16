package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Deletes an endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteEndpointCmd).Standalone()

		sagemaker_deleteEndpointCmd.Flags().String("endpoint-name", "", "The name of the endpoint that you want to delete.")
		sagemaker_deleteEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteEndpointCmd)
}
