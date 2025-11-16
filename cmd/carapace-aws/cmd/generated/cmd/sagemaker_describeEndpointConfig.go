package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeEndpointConfigCmd = &cobra.Command{
	Use:   "describe-endpoint-config",
	Short: "Returns the description of an endpoint configuration created using the `CreateEndpointConfig` API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeEndpointConfigCmd).Standalone()

	sagemaker_describeEndpointConfigCmd.Flags().String("endpoint-config-name", "", "The name of the endpoint configuration.")
	sagemaker_describeEndpointConfigCmd.MarkFlagRequired("endpoint-config-name")
	sagemakerCmd.AddCommand(sagemaker_describeEndpointConfigCmd)
}
