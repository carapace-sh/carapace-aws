package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getInvitationsCountCmd = &cobra.Command{
	Use:   "get-invitations-count",
	Short: "Retrieves the count of Amazon Macie membership invitations that were received by an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getInvitationsCountCmd).Standalone()

	macie2Cmd.AddCommand(macie2_getInvitationsCountCmd)
}
