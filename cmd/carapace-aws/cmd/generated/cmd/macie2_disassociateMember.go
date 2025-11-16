package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_disassociateMemberCmd = &cobra.Command{
	Use:   "disassociate-member",
	Short: "Disassociates an Amazon Macie administrator account from a member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_disassociateMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_disassociateMemberCmd).Standalone()

		macie2_disassociateMemberCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_disassociateMemberCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_disassociateMemberCmd)
}
