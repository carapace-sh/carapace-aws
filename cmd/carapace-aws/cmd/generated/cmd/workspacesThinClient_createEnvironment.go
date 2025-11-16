package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Creates an environment for your thin client devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_createEnvironmentCmd).Standalone()

	workspacesThinClient_createEnvironmentCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("desired-software-set-id", "", "The ID of the software set to apply.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("desktop-arn", "", "The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Secure Browser, or AppStream 2.0.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("desktop-endpoint", "", "The URL for the identity provider login (only for environments that use AppStream 2.0).")
	workspacesThinClient_createEnvironmentCmd.Flags().String("device-creation-tags", "", "A map of the key-value pairs of the tag or tags to assign to the newly created devices for this environment.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the Key Management Service key to use to encrypt the environment.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("maintenance-window", "", "A specification for a time window to apply software updates.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("name", "", "The name for the environment.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("software-set-update-mode", "", "An option to define which software updates to apply.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("software-set-update-schedule", "", "An option to define if software updates should be applied within a maintenance window.")
	workspacesThinClient_createEnvironmentCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	workspacesThinClient_createEnvironmentCmd.MarkFlagRequired("desktop-arn")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_createEnvironmentCmd)
}
