package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_deleteLayerVersionCmd = &cobra.Command{
	Use:   "delete-layer-version",
	Short: "Deletes a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_deleteLayerVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_deleteLayerVersionCmd).Standalone()

		lambda_deleteLayerVersionCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
		lambda_deleteLayerVersionCmd.Flags().String("version-number", "", "The version number.")
		lambda_deleteLayerVersionCmd.MarkFlagRequired("layer-name")
		lambda_deleteLayerVersionCmd.MarkFlagRequired("version-number")
	})
	lambdaCmd.AddCommand(lambda_deleteLayerVersionCmd)
}
