package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getSessionLoggerCmd = &cobra.Command{
	Use:   "get-session-logger",
	Short: "Gets details about a specific session logger resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getSessionLoggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_getSessionLoggerCmd).Standalone()

		workspacesWeb_getSessionLoggerCmd.Flags().String("session-logger-arn", "", "The ARN of the session logger.")
		workspacesWeb_getSessionLoggerCmd.MarkFlagRequired("session-logger-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_getSessionLoggerCmd)
}
