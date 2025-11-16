package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_publishLayerVersionCmd = &cobra.Command{
	Use:   "publish-layer-version",
	Short: "Creates an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) from a ZIP archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_publishLayerVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_publishLayerVersionCmd).Standalone()

		lambda_publishLayerVersionCmd.Flags().String("compatible-architectures", "", "A list of compatible [instruction set architectures](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html).")
		lambda_publishLayerVersionCmd.Flags().String("compatible-runtimes", "", "A list of compatible [function runtimes](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Used for filtering with [ListLayers]() and [ListLayerVersions]().")
		lambda_publishLayerVersionCmd.Flags().String("content", "", "The function layer archive.")
		lambda_publishLayerVersionCmd.Flags().String("description", "", "The description of the version.")
		lambda_publishLayerVersionCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
		lambda_publishLayerVersionCmd.Flags().String("license-info", "", "The layer's software license.")
		lambda_publishLayerVersionCmd.MarkFlagRequired("content")
		lambda_publishLayerVersionCmd.MarkFlagRequired("layer-name")
	})
	lambdaCmd.AddCommand(lambda_publishLayerVersionCmd)
}
