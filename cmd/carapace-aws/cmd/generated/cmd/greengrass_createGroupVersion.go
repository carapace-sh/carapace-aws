package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createGroupVersionCmd = &cobra.Command{
	Use:   "create-group-version",
	Short: "Creates a version of a group which has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createGroupVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createGroupVersionCmd).Standalone()

		greengrass_createGroupVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createGroupVersionCmd.Flags().String("connector-definition-version-arn", "", "The ARN of the connector definition version for this group.")
		greengrass_createGroupVersionCmd.Flags().String("core-definition-version-arn", "", "The ARN of the core definition version for this group.")
		greengrass_createGroupVersionCmd.Flags().String("device-definition-version-arn", "", "The ARN of the device definition version for this group.")
		greengrass_createGroupVersionCmd.Flags().String("function-definition-version-arn", "", "The ARN of the function definition version for this group.")
		greengrass_createGroupVersionCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_createGroupVersionCmd.Flags().String("logger-definition-version-arn", "", "The ARN of the logger definition version for this group.")
		greengrass_createGroupVersionCmd.Flags().String("resource-definition-version-arn", "", "The ARN of the resource definition version for this group.")
		greengrass_createGroupVersionCmd.Flags().String("subscription-definition-version-arn", "", "The ARN of the subscription definition version for this group.")
		greengrass_createGroupVersionCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_createGroupVersionCmd)
}
