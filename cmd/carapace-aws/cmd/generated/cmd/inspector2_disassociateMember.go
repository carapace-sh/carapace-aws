package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_disassociateMemberCmd = &cobra.Command{
	Use:   "disassociate-member",
	Short: "Disassociates a member account from an Amazon Inspector delegated administrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_disassociateMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_disassociateMemberCmd).Standalone()

		inspector2_disassociateMemberCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the member account to disassociate.")
		inspector2_disassociateMemberCmd.MarkFlagRequired("account-id")
	})
	inspector2Cmd.AddCommand(inspector2_disassociateMemberCmd)
}
