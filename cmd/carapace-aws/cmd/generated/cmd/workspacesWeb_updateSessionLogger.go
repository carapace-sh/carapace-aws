package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateSessionLoggerCmd = &cobra.Command{
	Use:   "update-session-logger",
	Short: "Updates the details of a session logger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateSessionLoggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_updateSessionLoggerCmd).Standalone()

		workspacesWeb_updateSessionLoggerCmd.Flags().String("display-name", "", "The updated display name.")
		workspacesWeb_updateSessionLoggerCmd.Flags().String("event-filter", "", "The updated eventFilter.")
		workspacesWeb_updateSessionLoggerCmd.Flags().String("log-configuration", "", "The updated logConfiguration.")
		workspacesWeb_updateSessionLoggerCmd.Flags().String("session-logger-arn", "", "The ARN of the session logger to update.")
		workspacesWeb_updateSessionLoggerCmd.MarkFlagRequired("session-logger-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_updateSessionLoggerCmd)
}
