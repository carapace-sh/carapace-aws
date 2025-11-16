package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listLayerVersionsCmd = &cobra.Command{
	Use:   "list-layer-versions",
	Short: "Lists the versions of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listLayerVersionsCmd).Standalone()

	lambda_listLayerVersionsCmd.Flags().String("compatible-architecture", "", "The compatible [instruction set architecture](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html).")
	lambda_listLayerVersionsCmd.Flags().String("compatible-runtime", "", "A runtime identifier.")
	lambda_listLayerVersionsCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
	lambda_listLayerVersionsCmd.Flags().String("marker", "", "A pagination token returned by a previous call.")
	lambda_listLayerVersionsCmd.Flags().String("max-items", "", "The maximum number of versions to return.")
	lambda_listLayerVersionsCmd.MarkFlagRequired("layer-name")
	lambdaCmd.AddCommand(lambda_listLayerVersionsCmd)
}
