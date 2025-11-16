package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateEndpointCmd = &cobra.Command{
	Use:   "update-endpoint",
	Short: "Deploys the `EndpointConfig` specified in the request to a new fleet of instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateEndpointCmd).Standalone()

		sagemaker_updateEndpointCmd.Flags().String("deployment-config", "", "The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.")
		sagemaker_updateEndpointCmd.Flags().String("endpoint-config-name", "", "The name of the new endpoint configuration.")
		sagemaker_updateEndpointCmd.Flags().String("endpoint-name", "", "The name of the endpoint whose configuration you want to update.")
		sagemaker_updateEndpointCmd.Flags().String("exclude-retained-variant-properties", "", "When you are updating endpoint resources with `RetainAllVariantProperties`, whose value is set to `true`, `ExcludeRetainedVariantProperties` specifies the list of type [VariantProperty](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VariantProperty.html) to override with the values provided by `EndpointConfig`.")
		sagemaker_updateEndpointCmd.Flags().Bool("no-retain-all-variant-properties", false, "When updating endpoint resources, enables or disables the retention of [variant properties](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VariantProperty.html), such as the instance count or the variant weight.")
		sagemaker_updateEndpointCmd.Flags().Bool("no-retain-deployment-config", false, "Specifies whether to reuse the last deployment configuration.")
		sagemaker_updateEndpointCmd.Flags().Bool("retain-all-variant-properties", false, "When updating endpoint resources, enables or disables the retention of [variant properties](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VariantProperty.html), such as the instance count or the variant weight.")
		sagemaker_updateEndpointCmd.Flags().Bool("retain-deployment-config", false, "Specifies whether to reuse the last deployment configuration.")
		sagemaker_updateEndpointCmd.MarkFlagRequired("endpoint-config-name")
		sagemaker_updateEndpointCmd.MarkFlagRequired("endpoint-name")
		sagemaker_updateEndpointCmd.Flag("no-retain-all-variant-properties").Hidden = true
		sagemaker_updateEndpointCmd.Flag("no-retain-deployment-config").Hidden = true
	})
	sagemakerCmd.AddCommand(sagemaker_updateEndpointCmd)
}
