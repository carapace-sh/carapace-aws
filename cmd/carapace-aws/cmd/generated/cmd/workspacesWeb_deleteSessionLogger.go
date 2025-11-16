package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteSessionLoggerCmd = &cobra.Command{
	Use:   "delete-session-logger",
	Short: "Deletes a session logger resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteSessionLoggerCmd).Standalone()

	workspacesWeb_deleteSessionLoggerCmd.Flags().String("session-logger-arn", "", "The ARN of the session logger.")
	workspacesWeb_deleteSessionLoggerCmd.MarkFlagRequired("session-logger-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteSessionLoggerCmd)
}
