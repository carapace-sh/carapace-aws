package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getInvitationsCountCmd = &cobra.Command{
	Use:   "get-invitations-count",
	Short: "We recommend using Organizations instead of Security Hub invitations to manage your member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getInvitationsCountCmd).Standalone()

	securityhubCmd.AddCommand(securityhub_getInvitationsCountCmd)
}
