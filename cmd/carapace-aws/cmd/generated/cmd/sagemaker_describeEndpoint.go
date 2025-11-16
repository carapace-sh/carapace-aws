package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeEndpointCmd = &cobra.Command{
	Use:   "describe-endpoint",
	Short: "Returns the description of an endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeEndpointCmd).Standalone()

		sagemaker_describeEndpointCmd.Flags().String("endpoint-name", "", "The name of the endpoint.")
		sagemaker_describeEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeEndpointCmd)
}
