package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateUserSettingsCmd = &cobra.Command{
	Use:   "update-user-settings",
	Short: "Updates the user settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateUserSettingsCmd).Standalone()

	workspacesWeb_updateUserSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("cookie-synchronization-configuration", "", "The configuration that specifies which cookies should be synchronized from the end user's local browser to the remote browser.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("copy-allowed", "", "Specifies whether the user can copy text from the streaming session to the local device.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("deep-link-allowed", "", "Specifies whether the user can use deep links that open automatically when connecting to a session.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("disconnect-timeout-in-minutes", "", "The amount of time that a streaming session remains active after users disconnect.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("download-allowed", "", "Specifies whether the user can download files from the streaming session to the local device.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("idle-disconnect-timeout-in-minutes", "", "The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the disconnect timeout interval begins.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("paste-allowed", "", "Specifies whether the user can paste text from the local device to the streaming session.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("print-allowed", "", "Specifies whether the user can print to the local device.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("toolbar-configuration", "", "The configuration of the toolbar.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("upload-allowed", "", "Specifies whether the user can upload files from the local device to the streaming session.")
	workspacesWeb_updateUserSettingsCmd.Flags().String("user-settings-arn", "", "The ARN of the user settings.")
	workspacesWeb_updateUserSettingsCmd.MarkFlagRequired("user-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updateUserSettingsCmd)
}
