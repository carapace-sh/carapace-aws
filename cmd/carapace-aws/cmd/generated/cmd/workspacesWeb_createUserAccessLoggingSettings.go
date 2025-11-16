package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "create-user-access-logging-settings",
	Short: "Creates a user access logging settings resource that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createUserAccessLoggingSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_createUserAccessLoggingSettingsCmd).Standalone()

		workspacesWeb_createUserAccessLoggingSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_createUserAccessLoggingSettingsCmd.Flags().String("kinesis-stream-arn", "", "The ARN of the Kinesis stream.")
		workspacesWeb_createUserAccessLoggingSettingsCmd.Flags().String("tags", "", "The tags to add to the user settings resource.")
		workspacesWeb_createUserAccessLoggingSettingsCmd.MarkFlagRequired("kinesis-stream-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_createUserAccessLoggingSettingsCmd)
}
