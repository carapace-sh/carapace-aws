package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_expireSessionCmd = &cobra.Command{
	Use:   "expire-session",
	Short: "Expires an active secure browser session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_expireSessionCmd).Standalone()

	workspacesWeb_expireSessionCmd.Flags().String("portal-id", "", "The ID of the web portal for the session.")
	workspacesWeb_expireSessionCmd.Flags().String("session-id", "", "The ID of the session to expire.")
	workspacesWeb_expireSessionCmd.MarkFlagRequired("portal-id")
	workspacesWeb_expireSessionCmd.MarkFlagRequired("session-id")
	workspacesWebCmd.AddCommand(workspacesWeb_expireSessionCmd)
}
