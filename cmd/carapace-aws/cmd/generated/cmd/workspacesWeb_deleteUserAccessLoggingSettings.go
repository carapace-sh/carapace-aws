package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "delete-user-access-logging-settings",
	Short: "Deletes user access logging settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteUserAccessLoggingSettingsCmd).Standalone()

	workspacesWeb_deleteUserAccessLoggingSettingsCmd.Flags().String("user-access-logging-settings-arn", "", "The ARN of the user access logging settings.")
	workspacesWeb_deleteUserAccessLoggingSettingsCmd.MarkFlagRequired("user-access-logging-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteUserAccessLoggingSettingsCmd)
}
