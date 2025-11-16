package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "associate-user-access-logging-settings",
	Short: "Associates a user access logging settings resource with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateUserAccessLoggingSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateUserAccessLoggingSettingsCmd).Standalone()

		workspacesWeb_associateUserAccessLoggingSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_associateUserAccessLoggingSettingsCmd.Flags().String("user-access-logging-settings-arn", "", "The ARN of the user access logging settings.")
		workspacesWeb_associateUserAccessLoggingSettingsCmd.MarkFlagRequired("portal-arn")
		workspacesWeb_associateUserAccessLoggingSettingsCmd.MarkFlagRequired("user-access-logging-settings-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateUserAccessLoggingSettingsCmd)
}
