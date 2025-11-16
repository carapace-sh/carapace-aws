package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_deleteResourceCmd = &cobra.Command{
	Use:   "delete-resource",
	Short: "Deletes the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_deleteResourceCmd).Standalone()

	cloudcontrol_deleteResourceCmd.Flags().String("client-token", "", "A unique identifier to ensure the idempotency of the resource request.")
	cloudcontrol_deleteResourceCmd.Flags().String("identifier", "", "The identifier for the resource.")
	cloudcontrol_deleteResourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation.")
	cloudcontrol_deleteResourceCmd.Flags().String("type-name", "", "The name of the resource type.")
	cloudcontrol_deleteResourceCmd.Flags().String("type-version-id", "", "For private resource types, the type version to use in this resource operation.")
	cloudcontrol_deleteResourceCmd.MarkFlagRequired("identifier")
	cloudcontrol_deleteResourceCmd.MarkFlagRequired("type-name")
	cloudcontrolCmd.AddCommand(cloudcontrol_deleteResourceCmd)
}
