package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Retrieves information for a specified session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_getSessionCmd).Standalone()

		wisdom_getSessionCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
		wisdom_getSessionCmd.Flags().String("session-id", "", "The identifier of the session.")
		wisdom_getSessionCmd.MarkFlagRequired("assistant-id")
		wisdom_getSessionCmd.MarkFlagRequired("session-id")
	})
	wisdomCmd.AddCommand(wisdom_getSessionCmd)
}
