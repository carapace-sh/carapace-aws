package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_getResourceCmd = &cobra.Command{
	Use:   "get-resource",
	Short: "Returns information about the current state of the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_getResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudcontrol_getResourceCmd).Standalone()

		cloudcontrol_getResourceCmd.Flags().String("identifier", "", "The identifier for the resource.")
		cloudcontrol_getResourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation.")
		cloudcontrol_getResourceCmd.Flags().String("type-name", "", "The name of the resource type.")
		cloudcontrol_getResourceCmd.Flags().String("type-version-id", "", "For private resource types, the type version to use in this resource operation.")
		cloudcontrol_getResourceCmd.MarkFlagRequired("identifier")
		cloudcontrol_getResourceCmd.MarkFlagRequired("type-name")
	})
	cloudcontrolCmd.AddCommand(cloudcontrol_getResourceCmd)
}
