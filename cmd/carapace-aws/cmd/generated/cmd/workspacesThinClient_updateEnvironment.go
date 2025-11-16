package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Updates an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_updateEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_updateEnvironmentCmd).Standalone()

		workspacesThinClient_updateEnvironmentCmd.Flags().String("desired-software-set-id", "", "The ID of the software set to apply.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("desktop-arn", "", "The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Secure Browser, or AppStream 2.0.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("desktop-endpoint", "", "The URL for the identity provider login (only for environments that use AppStream 2.0).")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("device-creation-tags", "", "A map of the key-value pairs of the tag or tags to assign to the newly created devices for this environment.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("id", "", "The ID of the environment to update.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("maintenance-window", "", "A specification for a time window to apply software updates.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("name", "", "The name of the environment to update.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("software-set-update-mode", "", "An option to define which software updates to apply.")
		workspacesThinClient_updateEnvironmentCmd.Flags().String("software-set-update-schedule", "", "An option to define if software updates should be applied within a maintenance window.")
		workspacesThinClient_updateEnvironmentCmd.MarkFlagRequired("id")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_updateEnvironmentCmd)
}
