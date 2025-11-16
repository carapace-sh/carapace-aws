package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Gets the full details of a previously created session, including the session status and configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getSessionCmd).Standalone()

		athena_getSessionCmd.Flags().String("session-id", "", "The session ID.")
		athena_getSessionCmd.MarkFlagRequired("session-id")
	})
	athenaCmd.AddCommand(athena_getSessionCmd)
}
