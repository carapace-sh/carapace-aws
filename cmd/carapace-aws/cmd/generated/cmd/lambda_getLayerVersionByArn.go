package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_getLayerVersionByArnCmd = &cobra.Command{
	Use:   "get-layer-version-by-arn",
	Short: "Returns information about a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html), with a link to download the layer archive that's valid for 10 minutes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_getLayerVersionByArnCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_getLayerVersionByArnCmd).Standalone()

		lambda_getLayerVersionByArnCmd.Flags().String("arn", "", "The ARN of the layer version.")
		lambda_getLayerVersionByArnCmd.MarkFlagRequired("arn")
	})
	lambdaCmd.AddCommand(lambda_getLayerVersionByArnCmd)
}
