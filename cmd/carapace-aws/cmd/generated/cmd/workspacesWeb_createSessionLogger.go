package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createSessionLoggerCmd = &cobra.Command{
	Use:   "create-session-logger",
	Short: "Creates a session logger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createSessionLoggerCmd).Standalone()

	workspacesWeb_createSessionLoggerCmd.Flags().String("additional-encryption-context", "", "The additional encryption context of the session logger.")
	workspacesWeb_createSessionLoggerCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_createSessionLoggerCmd.Flags().String("customer-managed-key", "", "The custom managed key of the session logger.")
	workspacesWeb_createSessionLoggerCmd.Flags().String("display-name", "", "The human-readable display name for the session logger resource.")
	workspacesWeb_createSessionLoggerCmd.Flags().String("event-filter", "", "The filter that specifies the events to monitor.")
	workspacesWeb_createSessionLoggerCmd.Flags().String("log-configuration", "", "The configuration that specifies where logs are delivered.")
	workspacesWeb_createSessionLoggerCmd.Flags().String("tags", "", "The tags to add to the session logger.")
	workspacesWeb_createSessionLoggerCmd.MarkFlagRequired("event-filter")
	workspacesWeb_createSessionLoggerCmd.MarkFlagRequired("log-configuration")
	workspacesWebCmd.AddCommand(workspacesWeb_createSessionLoggerCmd)
}
