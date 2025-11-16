package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getSessionStatusCmd = &cobra.Command{
	Use:   "get-session-status",
	Short: "Gets the current status of a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getSessionStatusCmd).Standalone()

	athena_getSessionStatusCmd.Flags().String("session-id", "", "The session ID.")
	athena_getSessionStatusCmd.MarkFlagRequired("session-id")
	athenaCmd.AddCommand(athena_getSessionStatusCmd)
}
