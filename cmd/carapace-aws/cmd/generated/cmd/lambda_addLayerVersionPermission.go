package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_addLayerVersionPermissionCmd = &cobra.Command{
	Use:   "add-layer-version-permission",
	Short: "Adds permissions to the resource-based policy of a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). Use this action to grant layer usage permission to other accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_addLayerVersionPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_addLayerVersionPermissionCmd).Standalone()

		lambda_addLayerVersionPermissionCmd.Flags().String("action", "", "The API action that grants access to the layer.")
		lambda_addLayerVersionPermissionCmd.Flags().String("layer-name", "", "The name or Amazon Resource Name (ARN) of the layer.")
		lambda_addLayerVersionPermissionCmd.Flags().String("organization-id", "", "With the principal set to `*`, grant permission to all accounts in the specified organization.")
		lambda_addLayerVersionPermissionCmd.Flags().String("principal", "", "An account ID, or `*` to grant layer usage permission to all accounts in an organization, or all Amazon Web Services accounts (if `organizationId` is not specified).")
		lambda_addLayerVersionPermissionCmd.Flags().String("revision-id", "", "Only update the policy if the revision ID matches the ID specified.")
		lambda_addLayerVersionPermissionCmd.Flags().String("statement-id", "", "An identifier that distinguishes the policy from others on the same layer version.")
		lambda_addLayerVersionPermissionCmd.Flags().String("version-number", "", "The version number.")
		lambda_addLayerVersionPermissionCmd.MarkFlagRequired("action")
		lambda_addLayerVersionPermissionCmd.MarkFlagRequired("layer-name")
		lambda_addLayerVersionPermissionCmd.MarkFlagRequired("principal")
		lambda_addLayerVersionPermissionCmd.MarkFlagRequired("statement-id")
		lambda_addLayerVersionPermissionCmd.MarkFlagRequired("version-number")
	})
	lambdaCmd.AddCommand(lambda_addLayerVersionPermissionCmd)
}
