package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateSessionLoggerCmd = &cobra.Command{
	Use:   "disassociate-session-logger",
	Short: "Disassociates a session logger from a portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateSessionLoggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_disassociateSessionLoggerCmd).Standalone()

		workspacesWeb_disassociateSessionLoggerCmd.Flags().String("portal-arn", "", "The ARN of the portal to disassociate from the a session logger.")
		workspacesWeb_disassociateSessionLoggerCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateSessionLoggerCmd)
}
