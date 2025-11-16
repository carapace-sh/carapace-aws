package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createEndpointConfigCmd = &cobra.Command{
	Use:   "create-endpoint-config",
	Short: "Creates an endpoint configuration that SageMaker hosting services uses to deploy models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createEndpointConfigCmd).Standalone()

	sagemaker_createEndpointConfigCmd.Flags().String("async-inference-config", "", "Specifies configuration for how an endpoint performs asynchronous inference.")
	sagemaker_createEndpointConfigCmd.Flags().String("data-capture-config", "", "")
	sagemaker_createEndpointConfigCmd.Flags().Bool("enable-network-isolation", false, "Sets whether all model containers deployed to the endpoint are isolated.")
	sagemaker_createEndpointConfigCmd.Flags().String("endpoint-config-name", "", "The name of the endpoint configuration.")
	sagemaker_createEndpointConfigCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform actions on your behalf.")
	sagemaker_createEndpointConfigCmd.Flags().String("explainer-config", "", "A member of `CreateEndpointConfig` that enables explainers.")
	sagemaker_createEndpointConfigCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.")
	sagemaker_createEndpointConfigCmd.Flags().Bool("no-enable-network-isolation", false, "Sets whether all model containers deployed to the endpoint are isolated.")
	sagemaker_createEndpointConfigCmd.Flags().String("production-variants", "", "An array of `ProductionVariant` objects, one for each model that you want to host at this endpoint.")
	sagemaker_createEndpointConfigCmd.Flags().String("shadow-production-variants", "", "An array of `ProductionVariant` objects, one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on `ProductionVariants`.")
	sagemaker_createEndpointConfigCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createEndpointConfigCmd.Flags().String("vpc-config", "", "")
	sagemaker_createEndpointConfigCmd.MarkFlagRequired("endpoint-config-name")
	sagemaker_createEndpointConfigCmd.Flag("no-enable-network-isolation").Hidden = true
	sagemaker_createEndpointConfigCmd.MarkFlagRequired("production-variants")
	sagemakerCmd.AddCommand(sagemaker_createEndpointConfigCmd)
}
