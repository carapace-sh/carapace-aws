package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getLayerVersionPolicyCmd = &cobra.Command{
	Use:   "get-layer-version-policy",
	Short: "Returns the permission policy for a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). For more information, see [AddLayerVersionPermission]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getLayerVersionPolicyCmd).Standalone()

	lambda_getLayerVersionPolicyCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
	lambda_getLayerVersionPolicyCmd.Flags().String("version-number", "", "The version number.")
	lambda_getLayerVersionPolicyCmd.MarkFlagRequired("layer-name")
	lambda_getLayerVersionPolicyCmd.MarkFlagRequired("version-number")
	lambdaCmd.AddCommand(lambda_getLayerVersionPolicyCmd)
}
