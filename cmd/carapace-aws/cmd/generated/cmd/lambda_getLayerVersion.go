package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getLayerVersionCmd = &cobra.Command{
	Use:   "get-layer-version",
	Short: "Returns information about a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html), with a link to download the layer archive that's valid for 10 minutes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getLayerVersionCmd).Standalone()

	lambda_getLayerVersionCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
	lambda_getLayerVersionCmd.Flags().String("version-number", "", "The version number.")
	lambda_getLayerVersionCmd.MarkFlagRequired("layer-name")
	lambda_getLayerVersionCmd.MarkFlagRequired("version-number")
	lambdaCmd.AddCommand(lambda_getLayerVersionCmd)
}
