package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_updateResourceCmd = &cobra.Command{
	Use:   "update-resource",
	Short: "Updates the specified property values in the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_updateResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudcontrol_updateResourceCmd).Standalone()

		cloudcontrol_updateResourceCmd.Flags().String("client-token", "", "A unique identifier to ensure the idempotency of the resource request.")
		cloudcontrol_updateResourceCmd.Flags().String("identifier", "", "The identifier for the resource.")
		cloudcontrol_updateResourceCmd.Flags().String("patch-document", "", "A JavaScript Object Notation (JSON) document listing the patch operations that represent the updates to apply to the current resource properties.")
		cloudcontrol_updateResourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation.")
		cloudcontrol_updateResourceCmd.Flags().String("type-name", "", "The name of the resource type.")
		cloudcontrol_updateResourceCmd.Flags().String("type-version-id", "", "For private resource types, the type version to use in this resource operation.")
		cloudcontrol_updateResourceCmd.MarkFlagRequired("identifier")
		cloudcontrol_updateResourceCmd.MarkFlagRequired("patch-document")
		cloudcontrol_updateResourceCmd.MarkFlagRequired("type-name")
	})
	cloudcontrolCmd.AddCommand(cloudcontrol_updateResourceCmd)
}
