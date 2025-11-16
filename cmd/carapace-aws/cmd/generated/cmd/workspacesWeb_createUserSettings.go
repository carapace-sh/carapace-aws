package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createUserSettingsCmd = &cobra.Command{
	Use:   "create-user-settings",
	Short: "Creates a user settings resource that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createUserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_createUserSettingsCmd).Standalone()

		workspacesWeb_createUserSettingsCmd.Flags().String("additional-encryption-context", "", "The additional encryption context of the user settings.")
		workspacesWeb_createUserSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_createUserSettingsCmd.Flags().String("cookie-synchronization-configuration", "", "The configuration that specifies which cookies should be synchronized from the end user's local browser to the remote browser.")
		workspacesWeb_createUserSettingsCmd.Flags().String("copy-allowed", "", "Specifies whether the user can copy text from the streaming session to the local device.")
		workspacesWeb_createUserSettingsCmd.Flags().String("customer-managed-key", "", "The customer managed key used to encrypt sensitive information in the user settings.")
		workspacesWeb_createUserSettingsCmd.Flags().String("deep-link-allowed", "", "Specifies whether the user can use deep links that open automatically when connecting to a session.")
		workspacesWeb_createUserSettingsCmd.Flags().String("disconnect-timeout-in-minutes", "", "The amount of time that a streaming session remains active after users disconnect.")
		workspacesWeb_createUserSettingsCmd.Flags().String("download-allowed", "", "Specifies whether the user can download files from the streaming session to the local device.")
		workspacesWeb_createUserSettingsCmd.Flags().String("idle-disconnect-timeout-in-minutes", "", "The amount of time that users can be idle (inactive) before they are disconnected from their streaming session and the disconnect timeout interval begins.")
		workspacesWeb_createUserSettingsCmd.Flags().String("paste-allowed", "", "Specifies whether the user can paste text from the local device to the streaming session.")
		workspacesWeb_createUserSettingsCmd.Flags().String("print-allowed", "", "Specifies whether the user can print to the local device.")
		workspacesWeb_createUserSettingsCmd.Flags().String("tags", "", "The tags to add to the user settings resource.")
		workspacesWeb_createUserSettingsCmd.Flags().String("toolbar-configuration", "", "The configuration of the toolbar.")
		workspacesWeb_createUserSettingsCmd.Flags().String("upload-allowed", "", "Specifies whether the user can upload files from the local device to the streaming session.")
		workspacesWeb_createUserSettingsCmd.MarkFlagRequired("copy-allowed")
		workspacesWeb_createUserSettingsCmd.MarkFlagRequired("download-allowed")
		workspacesWeb_createUserSettingsCmd.MarkFlagRequired("paste-allowed")
		workspacesWeb_createUserSettingsCmd.MarkFlagRequired("print-allowed")
		workspacesWeb_createUserSettingsCmd.MarkFlagRequired("upload-allowed")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_createUserSettingsCmd)
}
