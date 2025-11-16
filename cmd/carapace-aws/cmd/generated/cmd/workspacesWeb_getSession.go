package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Gets information for a secure browser session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_getSessionCmd).Standalone()

		workspacesWeb_getSessionCmd.Flags().String("portal-id", "", "The ID of the web portal for the session.")
		workspacesWeb_getSessionCmd.Flags().String("session-id", "", "The ID of the session.")
		workspacesWeb_getSessionCmd.MarkFlagRequired("portal-id")
		workspacesWeb_getSessionCmd.MarkFlagRequired("session-id")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_getSessionCmd)
}
