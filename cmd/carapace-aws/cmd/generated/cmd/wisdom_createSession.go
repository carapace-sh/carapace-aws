package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_createSessionCmd = &cobra.Command{
	Use:   "create-session",
	Short: "Creates a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_createSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_createSessionCmd).Standalone()

		wisdom_createSessionCmd.Flags().String("assistant-id", "", "The identifier of the Wisdom assistant.")
		wisdom_createSessionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		wisdom_createSessionCmd.Flags().String("description", "", "The description.")
		wisdom_createSessionCmd.Flags().String("name", "", "The name of the session.")
		wisdom_createSessionCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		wisdom_createSessionCmd.MarkFlagRequired("assistant-id")
		wisdom_createSessionCmd.MarkFlagRequired("name")
	})
	wisdomCmd.AddCommand(wisdom_createSessionCmd)
}
