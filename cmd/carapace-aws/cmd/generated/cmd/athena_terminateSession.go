package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_terminateSessionCmd = &cobra.Command{
	Use:   "terminate-session",
	Short: "Terminates an active session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_terminateSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_terminateSessionCmd).Standalone()

		athena_terminateSessionCmd.Flags().String("session-id", "", "The session ID.")
		athena_terminateSessionCmd.MarkFlagRequired("session-id")
	})
	athenaCmd.AddCommand(athena_terminateSessionCmd)
}
