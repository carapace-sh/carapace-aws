package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "update-user-access-logging-settings",
	Short: "Updates the user access logging settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateUserAccessLoggingSettingsCmd).Standalone()

	workspacesWeb_updateUserAccessLoggingSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_updateUserAccessLoggingSettingsCmd.Flags().String("kinesis-stream-arn", "", "The ARN of the Kinesis stream.")
	workspacesWeb_updateUserAccessLoggingSettingsCmd.Flags().String("user-access-logging-settings-arn", "", "The ARN of the user access logging settings.")
	workspacesWeb_updateUserAccessLoggingSettingsCmd.MarkFlagRequired("user-access-logging-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updateUserAccessLoggingSettingsCmd)
}
