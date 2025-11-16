package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "get-user-access-logging-settings",
	Short: "Gets user access logging settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getUserAccessLoggingSettingsCmd).Standalone()

	workspacesWeb_getUserAccessLoggingSettingsCmd.Flags().String("user-access-logging-settings-arn", "", "The ARN of the user access logging settings.")
	workspacesWeb_getUserAccessLoggingSettingsCmd.MarkFlagRequired("user-access-logging-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getUserAccessLoggingSettingsCmd)
}
