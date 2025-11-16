package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_removeLayerVersionPermissionCmd = &cobra.Command{
	Use:   "remove-layer-version-permission",
	Short: "Removes a statement from the permissions policy for a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). For more information, see [AddLayerVersionPermission]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_removeLayerVersionPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_removeLayerVersionPermissionCmd).Standalone()

		lambda_removeLayerVersionPermissionCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
		lambda_removeLayerVersionPermissionCmd.Flags().String("revision-id", "", "Only update the policy if the revision ID matches the ID specified.")
		lambda_removeLayerVersionPermissionCmd.Flags().String("statement-id", "", "The identifier that was specified when the statement was added.")
		lambda_removeLayerVersionPermissionCmd.Flags().String("version-number", "", "The version number.")
		lambda_removeLayerVersionPermissionCmd.MarkFlagRequired("layer-name")
		lambda_removeLayerVersionPermissionCmd.MarkFlagRequired("statement-id")
		lambda_removeLayerVersionPermissionCmd.MarkFlagRequired("version-number")
	})
	lambdaCmd.AddCommand(lambda_removeLayerVersionPermissionCmd)
}
