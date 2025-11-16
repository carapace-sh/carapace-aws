package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateSessionLoggerCmd = &cobra.Command{
	Use:   "associate-session-logger",
	Short: "Associates a session logger with a portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateSessionLoggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateSessionLoggerCmd).Standalone()

		workspacesWeb_associateSessionLoggerCmd.Flags().String("portal-arn", "", "The ARN of the portal to associate to the session logger ARN.")
		workspacesWeb_associateSessionLoggerCmd.Flags().String("session-logger-arn", "", "The ARN of the session logger to associate to the portal ARN.")
		workspacesWeb_associateSessionLoggerCmd.MarkFlagRequired("portal-arn")
		workspacesWeb_associateSessionLoggerCmd.MarkFlagRequired("session-logger-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateSessionLoggerCmd)
}
