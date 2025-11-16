package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_createResourceCmd = &cobra.Command{
	Use:   "create-resource",
	Short: "Creates the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_createResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudcontrol_createResourceCmd).Standalone()

		cloudcontrol_createResourceCmd.Flags().String("client-token", "", "A unique identifier to ensure the idempotency of the resource request.")
		cloudcontrol_createResourceCmd.Flags().String("desired-state", "", "Structured data format representing the desired state of the resource, consisting of that resource's properties and their desired values.")
		cloudcontrol_createResourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation.")
		cloudcontrol_createResourceCmd.Flags().String("type-name", "", "The name of the resource type.")
		cloudcontrol_createResourceCmd.Flags().String("type-version-id", "", "For private resource types, the type version to use in this resource operation.")
		cloudcontrol_createResourceCmd.MarkFlagRequired("desired-state")
		cloudcontrol_createResourceCmd.MarkFlagRequired("type-name")
	})
	cloudcontrolCmd.AddCommand(cloudcontrol_createResourceCmd)
}
