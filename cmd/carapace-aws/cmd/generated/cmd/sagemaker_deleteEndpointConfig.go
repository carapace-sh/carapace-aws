package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteEndpointConfigCmd = &cobra.Command{
	Use:   "delete-endpoint-config",
	Short: "Deletes an endpoint configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteEndpointConfigCmd).Standalone()

	sagemaker_deleteEndpointConfigCmd.Flags().String("endpoint-config-name", "", "The name of the endpoint configuration that you want to delete.")
	sagemaker_deleteEndpointConfigCmd.MarkFlagRequired("endpoint-config-name")
	sagemakerCmd.AddCommand(sagemaker_deleteEndpointConfigCmd)
}
