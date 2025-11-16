package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createEndpointCmd = &cobra.Command{
	Use:   "create-endpoint",
	Short: "Creates an endpoint using the endpoint configuration specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createEndpointCmd).Standalone()

	sagemaker_createEndpointCmd.Flags().String("deployment-config", "", "")
	sagemaker_createEndpointCmd.Flags().String("endpoint-config-name", "", "The name of an endpoint configuration.")
	sagemaker_createEndpointCmd.Flags().String("endpoint-name", "", "The name of the endpoint.The name must be unique within an Amazon Web Services Region in your Amazon Web Services account.")
	sagemaker_createEndpointCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createEndpointCmd.MarkFlagRequired("endpoint-config-name")
	sagemaker_createEndpointCmd.MarkFlagRequired("endpoint-name")
	sagemakerCmd.AddCommand(sagemaker_createEndpointCmd)
}
